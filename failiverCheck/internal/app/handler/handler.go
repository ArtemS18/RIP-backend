package handler

import (
	"failiverCheck/internal/app/models"
	"failiverCheck/internal/app/repository"

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
	r.PUT("/components/:id", h.UpdateComponent)
	r.POST("/components/", h.CreateComponent)
	r.POST("/components/:id/system_calc/", h.AddComponentInSystemCalc)
	r.DELETE("/components/:id", h.DeleteComponent)
	// r.GET("/availability_calc/:id", h.GetSystemCalc)
	// r.POST("/components", h.AddComponentInSystemCalc)
	// r.POST("/availability_calc", h.DeleteSystemCalc)
}

func (h *Handler) errorHandler(ctx *gin.Context, errorCode int, err error) {
	log.Error(err.Error())
	ctx.JSON(errorCode, gin.H{
		"status":      "error",
		"description": err.Error(),
	})

}

func (h *Handler) successHandler(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, models.OKResponse{
		Status: "ok",
		Data:   data,
	})

}
