package http

import (
	"failiverCheck/internal/app/config"
	"failiverCheck/internal/app/repository/minio"
	"failiverCheck/internal/app/repository/postgres"
	"failiverCheck/internal/app/usecase"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Postgres *postgres.Postgres
	Minio    *minio.Minio
	UseCase  *usecase.UseCase

	Config *config.Config
}

func NewHandler(pg *postgres.Postgres, m *minio.Minio, uc *usecase.UseCase, c *config.Config) *Handler {
	return &Handler{pg, m, uc, c}
}

func (h *Handler) RegisterHandlers(e *gin.Engine) {

	public := e.Group("/api")
	h.RegisterSwaggerHandlers(e)
	h.RegisterUserHandlers(public)
	h.RegisterComponentHandlers(public)
	h.RegisterSystemCalcHandlers(public)
	h.RegisterSysteemCalcToComponentsHandlers(public)

}

func (h *Handler) RegisterSwaggerHandlers(e *gin.Engine) {
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.PersistAuthorization(true)))
}

func (h *Handler) RegisterComponentHandlers(router *gin.RouterGroup) {
	// public endpoints
	router.GET("/components/", h.GetComponents)
	router.GET("/components/:id", h.GetComponent)
	// moderator endpoints
	protected := router.Group("", h.ModeratorValidateMiddleware())
	protected.PUT("/components/:id", h.UpdateComponent)
	protected.POST("/components/", h.CreateComponent)
	protected.DELETE("/components/:id", h.DeleteComponent)
	protected.POST("/components/:id/img", h.UpdateComponentImg)
	// users endpoints
	users := router.Group("", h.AuthoMiddleware())
	users.POST("/components/:id/system_calc/", h.AddComponentInSystemCalc)

}

func (h *Handler) RegisterSystemCalcHandlers(router *gin.RouterGroup) {

	users := router.Group("", h.AuthoMiddleware())
	users.GET("/system_calcs/", h.GetSystemCalcList)
	users.GET("/system_calcs/:id", h.SystemCalcAccessMiddleware(), h.GetSystemCalc)
	users.GET("/system_calcs/my_bucket", h.GetSystemCalcBucket)
	users.PUT("/system_calcs/:id", h.SystemCalcAccessMiddleware(), h.UpdateSystemCalc)
	users.PUT("/system_calcs/:id/status_formed", h.SystemCalcAccessMiddleware(), h.UpdateSystemCalcStatusToFormed)
	users.DELETE("/system_calcs/:id", h.SystemCalcAccessMiddleware(), h.DeleteSystemCalc)

	protected := router.Group("", h.ModeratorValidateMiddleware())
	protected.PUT("/system_calcs/:id/status", h.UpdateSystemCalcStatusModerator)
}

func (h *Handler) RegisterUserHandlers(router *gin.RouterGroup) {
	router.POST("/users/register", h.RegisterUser)
	router.POST("/users/auth", h.AuthUser)

	// users endpoints
	users := router.Group("", h.AuthoMiddleware())
	users.GET("/users/me", h.GetUser)
	users.PUT("/users/me", h.UpdateUser)
	users.POST("/users/logout", h.LogoutUser)
}
func (h *Handler) RegisterSysteemCalcToComponentsHandlers(router *gin.RouterGroup) {
	users := router.Group("", h.AuthoMiddleware())
	users.DELETE("/system_calcs_to_components/", h.DeleteComponentsToSystemCalc)
	users.PUT("/system_calcs_to_components/", h.UpdateComponentsToSystemCalc)
}
