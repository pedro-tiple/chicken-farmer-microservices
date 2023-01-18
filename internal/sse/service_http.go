package sse

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IController interface {
	SubscribeToFarmer(
		ctx context.Context, farmerID uuid.UUID,
	) (chan string, error)
}

type HTTPService struct {
	logger     *zap.SugaredLogger
	controller IController
	httpRouter chi.Router
}

func ProvideHTTPService(
	logger *zap.SugaredLogger,
	controller IController,
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
	}

	httpRouter.Get("/event-feed", service.HandleStreamConnect)

	return &service, nil
}

func (s *HTTPService) ListenAndServe(
	ctx context.Context, address string,
) error {
	return http.ListenAndServe(address, s.httpRouter)
}

func (s *HTTPService) HandleStreamConnect(
	w http.ResponseWriter, r *http.Request,
) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	// TODO parse request header for auth JWT and store farmerID in context
	farmerID := uuid.MustParse("65e4d8ff-8766-48a7-bfcd-7160d149a319")

	ctx := r.Context()
	subscription, err := s.controller.SubscribeToFarmer(ctx, farmerID)
	if err != nil {
		s.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encoding to base64 so the SSE implementation stops messing up the json string.
	// return base64.StdEncoding.EncodeToString(message.Payload), true

	// Set SSE headers.
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// TODO limit CORS

	// Write something just to open the connection
	if _, err := fmt.Fprintf(w, `data: {"status": "open"}\n\n`); err != nil {
		s.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	flusher.Flush()

OuterLoop:
	for {
		select {
		case <-ctx.Done():
			break OuterLoop

		case message := <-subscription:
			if _, err := fmt.Fprintf(w, "data: %s\n\n", message); err != nil {
				s.logger.Error(err.Error())
				continue
			}

			flusher.Flush()
		}
	}
}
