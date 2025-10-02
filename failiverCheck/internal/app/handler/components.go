package handler

import (
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/dto"
	"failiverCheck/internal/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetComponent(ctx *gin.Context) {
	id := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	log.Info(id)
	component, err := h.Repository.GetComponentById(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}

	h.successHandler(ctx, http.StatusOK, component)
}

func (h *Handler) GetComponents(ctx *gin.Context) {
	var components []ds.Component
	var err error
	searchQuery := ctx.Query("search")

	log.Info(searchQuery)
	if searchQuery == "" {
		components, err = h.Repository.GetComponents()
		if err != nil {
			h.errorHandler(ctx, http.StatusBadRequest, err)
			return
		}
	} else {
		components, err = h.Repository.GetComponentsByTitle(searchQuery)
		if err != nil {
			h.errorHandler(ctx, http.StatusNotFound, err)
			return
		}
	}
	h.successHandler(ctx, http.StatusOK, models.ComponentsRes{Components: components})
}

func (h *Handler) UpdateComponent(ctx *gin.Context) {
	id := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	var update dto.UpdateComponentDTO
	if err := ctx.BindJSON(&update); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	component, err := h.Repository.UpdateComponentById(uint(id), update)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}

	h.successHandler(ctx, http.StatusOK, component)
}

func (h *Handler) CreateComponent(ctx *gin.Context) {
	var create dto.CreateComponentDTO
	if err := ctx.BindJSON(&create); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	validate := validator.New()
	if err := validate.Struct(create); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	component, err := h.Repository.CreateComponent(create)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, http.StatusOK, component)
}

func (h *Handler) DeleteComponent(ctx *gin.Context) {
	id := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	component, err := h.Repository.GetComponentById(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}
	imgUrl := component.Img

	if err := h.Repository.DeletedComponentById(uint(id)); err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}
	if err := h.Repository.DeleteComponentImg(ctx, &imgUrl); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (h *Handler) AddComponentInSystemCalc(ctx *gin.Context) {
	var err error
	id := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	var userId uint = 1
	err = h.Repository.AddComponentInSystemCalc(uint(id), userId)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (h *Handler) UpdateComponentImg(ctx *gin.Context) {
	id := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	contentType, fileSize := h.getFileHeaders(ctx)
	if ctx.IsAborted() {
		return
	}
	location, err := h.Repository.UploadComponentImg(ctx, dto.ComponentImgCreateDTO{
		File:        ctx.Request.Body,
		FilePath:    "img/",
		FileSize:    fileSize,
		ContentType: contentType,
	})
	logrus.Info(location)
	if err != nil {
		h.errorHandler(ctx, 404, err)
		return
	}
	component, err := h.Repository.GetComponentById(id)
	if err != nil {
		h.errorHandler(ctx, 404, err)
		return
	}
	if component.Img != "" {
		if err = h.Repository.DeleteComponentImg(ctx, &component.Img); err != nil {
			logrus.Error(err)
		}
	}
	h.Repository.UpdateComponentById(uint(id), dto.UpdateComponentDTO{Img: &location})
	h.successHandler(ctx, 200, map[string]string{"img_url": location})
}
