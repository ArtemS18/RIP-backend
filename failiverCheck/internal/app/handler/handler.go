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
	api := e.Group("/api")
	components := api.Group("/components", AuthMiddleware())
	components.GET("/", h.GetComponents)
	components.GET("/:id", h.GetComponent)
	components.PUT("/:id", h.UpdateComponent)
	components.POST("/", h.CreateComponent)
	components.POST("/:id/system_calc/", h.AddComponentInSystemCalc)
	components.DELETE("/:id", h.DeleteComponent)
	components.POST("/:id/img", h.UpdateComponentImg)

	systemCacl := api.Group("/system_calcs", AuthMiddleware())
	systemCacl.GET("/", h.GetSystemCalcList)
	systemCacl.GET("/:id", h.GetSystemCalc)
	systemCacl.GET("/my_bucket", h.GetSystemCalcBucket)
	systemCacl.PUT("/:id", h.UpdateSystemCalc)
	systemCacl.PUT("/:id/status_formed", h.UpdateSystemCalcStatusToFormed)
	systemCacl.PUT("/:id/status", h.UpdateSystemCalcStatusModerator)
	systemCacl.DELETE("/:id", h.DeleteSystemCalc)

	systemCaclToComp := api.Group("/system_calcs_to_components", AuthMiddleware())
	systemCaclToComp.DELETE("/", h.DeleteComponentsToSystemCac)
	systemCaclToComp.PUT("/", h.UpdateComponentsToSystemCac)

	users := api.Group("/users", AuthMiddleware())
	users.POST("/register", h.RegisterUser)
	users.POST("/auth", h.AuthUser)
	users.GET("/me", h.GetUser)
	users.POST("/loguout", h.LogoutUser)
	// r.GET("/availability_calc/:id", h.GetSystemCalc)
	// r.POST("/components", h.AddComponentInSystemCalc)
	// r.POST("/availability_calc", h.DeleteSystemCalc)
}
