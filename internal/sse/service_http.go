package sse

import (
	"chicken-farmer/backend/internal/pkg/event"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	watermillHTTP "github.com/ThreeDotsLabs/watermill-http/v2/pkg/http"
	messagePkg "github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IController interface {
	RegisterFarmer(ctx context.Context, farmerID uuid.UUID) error
}

type HTTPService struct {
	logger     *zap.SugaredLogger
	controller IController
	subscriber messagePkg.Subscriber
	httpRouter chi.Router
}

func ProvideWatermillService(
	logger *zap.SugaredLogger,
	controller IController,
	subscriber messagePkg.Subscriber,
) (*HTTPService, error) {
	httpRouter := chi.NewRouter()
	httpRouter.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{"https://*", "http://*"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: false,
				MaxAge:           300,
			},
		),
	)

	// TODO maybe send initial farm state here?

	service := HTTPService{
		logger:     logger,
		controller: controller,
		httpRouter: httpRouter,
		subscriber: subscriber,
	}

	httpRouter.Get("/event-feed", service.HandleStreamConnect)

	return &service, nil
}

func (s *HTTPService) HandleStreamConnect(
	w http.ResponseWriter, r *http.Request,
) {
	// TODO parse request header for auth JWT and store farmerID in context
	farmerID := uuid.MustParse("65e4d8ff-8766-48a7-bfcd-7160d149a319")

	sseRouter, err := watermillHTTP.NewSSERouter(
		watermillHTTP.SSERouterConfig{
			UpstreamSubscriber: s.subscriber,
			ErrorHandler:       watermillHTTP.DefaultErrorHandler,
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		s.logger.Error(err.Error())
		return
	}

	if err := s.controller.RegisterFarmer(r.Context(), farmerID); err != nil {
		s.logger.Error(err.Error())
		return
	}

	// This needs to be outside the go func call otherwise it will terminate early.
	handler := sseRouter.AddHandler(
		fmt.Sprintf(event.UserEventsTopic, farmerID), streamAdapter{},
	)
	go handler(w, r)

	if err := sseRouter.Run(context.Background()); err != nil {
		s.logger.Error(err.Error())
		return
	}
}

func (s *HTTPService) ListenAndServe(
	ctx context.Context, address string,
) error {
	return http.ListenAndServe(address, s.httpRouter)
}

type streamAdapter struct{}

func (f streamAdapter) InitialStreamResponse(
	w http.ResponseWriter, r *http.Request,
) (any, bool) {
	return "accepted", true
}

func (f streamAdapter) NextStreamResponse(
	r *http.Request, message *messagePkg.Message,
) (any, bool) {
	// Encoding to base64 so the SSE implementation stops messing up the json string.
	return base64.StdEncoding.EncodeToString(message.Payload), true
}
