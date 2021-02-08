package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-encurtador-url/domain"
)

type handler struct {
	urlService domain.URLService
	baseUrl    string
	port       *int
}

func NewHandler(urlService domain.URLService, baseUrl string, port *int) http.Handler {
	handler := &handler{urlService, baseUrl, port}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger(), handler.recovery())
	router.GET("/r/:short", handler.redirect)
	router.POST("/api/encurtar", handler.shorten)
	router.GET("/api/stats/:short", handler.visualize)

	return router
}

func (h *handler) recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recovered := recover(); recovered != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
