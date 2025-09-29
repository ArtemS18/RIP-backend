package handler

import (
	"failiverCheck/internal/app/repository"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Repository *repository.Repository
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{r}
}

func (h *Handler) RegisterHandlers(r *gin.Engine) {
	r.GET("/components", h.GetComponents)
	r.GET("/components/:id", h.GetComponent)
	r.GET("/availability_calc/:id", h.GetSystemCalc)
	r.POST("/components", h.AddComponentInSystemCalc)
	r.POST("/availability_calc", h.DeleteSystemCalc)
}

func (h *Handler) RegisterStatic(r *gin.Engine, path string) {
	r.LoadHTMLGlob(fmt.Sprintf("%s/*", path))
	r.Static("/static", "./resources")
}

func (h *Handler) errorHandler(ctx *gin.Context, errorCode int, err error) {
	log.Error(err.Error())
	ctx.JSON(errorCode, gin.H{
		"status":      "error",
		"description": err.Error(),
	})

}
