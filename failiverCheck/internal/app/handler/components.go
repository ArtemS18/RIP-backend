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
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invaid param id=%s", idStr))
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
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invaid param id=%s", idStr))
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
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invaid param id=%s", idStr))
		return
	}

	if err = h.Repository.DeletedComponentById(uint(id)); err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (h *Handler) AddComponentInSystemCalc(ctx *gin.Context) {
	var err error
	strId := ctx.Param("id")
	var userId uint = 1
	componentId, err := strconv.Atoi(strId)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invaid param id=%s", strId))
		return
	}
	err = h.Repository.AddComponentInSystemCalc(uint(componentId), userId)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
