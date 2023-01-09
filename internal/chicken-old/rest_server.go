package chicken_old

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RESTServer struct {
	Root   *gin.RouterGroup
	Engine API
	Logger zap.SugaredLogger
}

func (server RESTServer) SetupHandlers() {
	server.Root.GET("/chickens/{chickenID}", server.getChickenHandler)
	server.Root.POST("/chickens/new", server.newChickenHandler)
}

func (server RESTServer) getChickenHandler(c *gin.Context) {
	var queryParams struct {
		ChickenID string `form:"chickenID" binding:"required"`
	}
	if err := c.ShouldBindQuery(queryParams); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	chicken, err := server.Engine.GetChicken(
		c.Request.Context(), queryParams.ChickenID,
	)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"proto-old": chicken,
	})
}

func (server RESTServer) newChickenHandler(c *gin.Context) {
	var body NewChickenRequest
	if err := c.ShouldBind(body); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	chickenID, err := server.Engine.NewChicken(
		c.Request.Context(), body.FarmID,
	)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, NewChickenResult{
		ChickenID: chickenID,
	})
}
