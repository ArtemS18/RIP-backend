package handler

import (
	"failiverCheck/internal/app/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repository *repository.Repository
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{r}
}

func (h *Handler) RegisterHandlers(e *gin.Engine) {
	r := e.Group("/api")
	r.GET("/components", h.GetComponents)
	r.GET("/components/:id", h.GetComponent)
	r.PUT("/components/:id", h.UpdateComponent)
	r.POST("/components/", h.CreateComponent)
	r.POST("/components/:id/system_calc/", h.AddComponentInSystemCalc)
	r.DELETE("/components/:id", h.DeleteComponent)
	r.POST("/components/:id/img", h.UpdateComponentImg)

	r.GET("/system_calculations", h.GetSystemCalcList)
	r.GET("/system_calculations/:id", h.GetSystemCalc)
	r.GET("/system_calculations/my_bucket", h.GetSystemCalcBucket)
	r.PUT("/system_calculations/:id", h.UpdateSystemCalc)
	r.PUT("/system_calculations/:id/status_formed", h.UpdateSystemCalcStatusToFormed)
	r.PUT("/system_calculations/:id/status", h.UpdateSystemCalcStatusModerator)
	// r.GET("/availability_calc/:id", h.GetSystemCalc)
	// r.POST("/components", h.AddComponentInSystemCalc)
	// r.POST("/availability_calc", h.DeleteSystemCalc)
}
