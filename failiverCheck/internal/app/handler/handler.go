package handler

import (
	"failiverCheck/internal/app/config"
	"failiverCheck/internal/app/repository/minio"
	"failiverCheck/internal/app/repository/postgres"
	"failiverCheck/internal/app/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Postgres *postgres.Postgres
	Minio    *minio.Minio
	UseCase  *usecase.UseCase
	Config   *config.Config
}

func NewHandler(pg *postgres.Postgres, m *minio.Minio, uc *usecase.UseCase, c *config.Config) *Handler {
	return &Handler{pg, m, uc, c}
}

func (h *Handler) RegisterHandlers(e *gin.Engine) {
	api := e.Group("/api")
	components := api.Group("/components")
	components.GET("/", h.GetComponents)
	components.GET("/:id", h.GetComponent)
	components.PUT("/:id", h.UpdateComponent)
	components.POST("/", h.CreateComponent)
	components.POST("/:id/system_calc/", h.AddComponentInSystemCalc)
	components.DELETE("/:id", h.DeleteComponent)
	components.POST("/:id/img", h.UpdateComponentImg)

	systemCacl := api.Group("/system_calcs", h.AuthMiddleware(UserRole))
	systemCacl.GET("/", h.GetSystemCalcList)
	systemCacl.GET("/:id", h.GetSystemCalc)
	systemCacl.GET("/my_bucket", h.GetSystemCalcBucket)
	systemCacl.PUT("/:id", h.UpdateSystemCalc)
	systemCacl.PUT("/:id/status_formed", h.UpdateSystemCalcStatusToFormed)
	systemCacl.PUT("/:id/status", h.UpdateSystemCalcStatusModerator)
	systemCacl.DELETE("/:id", h.DeleteSystemCalc)

	systemCaclToComp := api.Group("/system_calcs_to_components")
	systemCaclToComp.DELETE("/", h.DeleteComponentsToSystemCac)
	systemCaclToComp.PUT("/", h.UpdateComponentsToSystemCac)

	users := api.Group("/users")
	users.POST("/register", h.RegisterUser)
	users.POST("/auth", h.AuthUser)
	users.GET("/me", h.AuthMiddleware(UserRole), h.GetUser)
	users.PUT("/me", h.AuthMiddleware(UserRole), h.UpdateUser)
	users.POST("/logout", h.LogoutUser)
	// r.GET("/availability_calc/:id", h.GetSystemCalc)
	// r.POST("/components", h.AddComponentInSystemCalc)
	// r.POST("/availability_calc", h.DeleteSystemCalc)
}
