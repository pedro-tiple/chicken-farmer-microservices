package sse

import (
	"context"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IController interface {
	SubscribeToFarmer(
		ctx context.Context, farmerID uuid.UUID,
	) (chan sse.Event, error)
}

type HTTPService struct {
	logger     *zap.SugaredLogger
	controller IController
	httpRouter *gin.Engine
}

func ProvideHTTPService(
	logger *zap.SugaredLogger,
	controller IController,
) (*HTTPService, error) {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language"}
	router.Use(cors.New(config))

	// TODO maybe send initial farm state here?

	service := HTTPService{
		logger:     logger,
		controller: controller,
		httpRouter: router,
	}

	router.GET("/event-feed", headersMiddleware, service.HandleStreamConnect)

	return &service, nil
}

func (s *HTTPService) ListenAndServe(
	ctx context.Context, address string,
) error {
	return s.httpRouter.Run(address)
}

func (s *HTTPService) HandleStreamConnect(c *gin.Context) {
	// TODO parse request header for auth JWT and store farmerID in context
	farmerID := uuid.MustParse("65e4d8ff-8766-48a7-bfcd-7160d149a319")

	ctx := c.Request.Context()
	subscription, err := s.controller.SubscribeToFarmer(ctx, farmerID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Stream(
		func(w io.Writer) bool {
			select {
			case <-ctx.Done():
				return false

			case message := <-subscription:
				c.SSEvent(message.Event, message.Data)
				return true
			}
		},
	)
}

func headersMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", sse.ContentType)
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	c.Next()
}
