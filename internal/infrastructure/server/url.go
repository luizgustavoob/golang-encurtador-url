package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-encurtador-url/domain"
	"github.com/golang-encurtador-url/domain/logger"
)

func (h *handler) shorten(c *gin.Context) {
	u, created, err := h.urlService.FindOrCreateURL(h.extractURL(c.Request))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var status int
	if created {
		status = http.StatusCreated
	} else {
		status = http.StatusOK
	}

	shortURL := fmt.Sprintf("%s/r/%s", h.baseUrl, u.ID)

	c.Writer.WriteHeader(status)
	c.Header("Location", shortURL)
	c.Header("Link", fmt.Sprintf("<%s/api/stats/%s>; rel=\"stats\"", h.baseUrl, u.ID))

	logger.Logar("URL %s encurtada com sucesso", u.Destination, shortURL)
}

func (h *handler) visualize(c *gin.Context) {
	h.findURLAndExecute(c, func(c *gin.Context, url *domain.Url) {
		c.JSONP(http.StatusOK, h.urlService.GetStatistics(url))
	})
}

func (h *handler) redirect(c *gin.Context) {
	h.findURLAndExecute(c, func(c *gin.Context, url *domain.Url) {
		c.Redirect(http.StatusMovedPermanently, url.Destination)
		h.urlService.AddStatistics(url)
	})
}

func (h *handler) findURLAndExecute(c *gin.Context, executor func(*gin.Context, *domain.Url)) {
	id := c.Param("short")
	if u := h.urlService.Find(id); u != nil {
		executor(c, u)
	} else {
		c.Writer.WriteHeader(http.StatusNotFound)
	}
}

func (h *handler) extractURL(r *http.Request) string {
	u := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(u)
	return string(u)
}
