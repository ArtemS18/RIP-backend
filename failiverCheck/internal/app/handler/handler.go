package handler

import (
	"failiverCheck/internal/app/models"
	"failiverCheck/internal/app/repository"
	"fmt"
	"strconv"

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
	r.POST("/components/:id/img", h.UpdateComponentImg)
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
	ctx.Abort()

}

func (h *Handler) successHandler(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, models.OKResponse{
		Status: "ok",
		Data:   data,
	})

}

func (h *Handler) getIntParam(ctx *gin.Context, param string) int {
	raw := ctx.Param(param)
	paramInt, err := strconv.Atoi(raw)
	if err != nil {
		h.errorHandler(ctx, 400, fmt.Errorf("invaid param id=%s", raw))
		return 0
	}
	return paramInt

}
