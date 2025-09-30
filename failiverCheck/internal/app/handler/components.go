package handler

import (
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/dto"
	"failiverCheck/internal/app/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	if err := h.Repository.DeletedComponentById(uint(id)); err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
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
	contentType := ctx.Request.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	contentLengthStr := ctx.Request.Header.Get("Content-Length")
	if contentLengthStr == "" {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("content-Length header is required"))
		return
	}
	fileSize, err := strconv.ParseInt(contentLengthStr, 10, 64)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid Content-Length header"))
		return
	}
	_, err = h.Repository.UploadComponentImg(ctx, ctx.Request.Body, fileSize, contentType)
	if err != nil {
		h.errorHandler(ctx, 500, err)
		return
	}
}
