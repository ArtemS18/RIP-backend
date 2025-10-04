package handler

import (
	"failiverCheck/internal/app/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *Handler) UpdateComponentsToSystemCac(ctx *gin.Context) {
	var update dto.UpdateComponentToSystemCalcDTO
	h.validateFields(ctx, &update)
	if ctx.IsAborted() {
		return
	}
	orm, err := h.Repository.UpdateComponentsToSystemCalc(update)
	new := dto.ToComponentsToSystemCalcDTO(orm)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}

	h.successHandler(ctx, http.StatusOK, new)
}

func (h *Handler) DeleteComponentsToSystemCac(ctx *gin.Context) {
	var ids dto.ComponentToSystemCalcByIdDTO
	if err := ctx.BindJSON(&ids); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	validate := validator.New()
	if err := validate.Struct(ids); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	err := h.Repository.DeleteComponentsToSystemCalc(ids)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}

	h.successHandler(ctx, http.StatusOK, nil)
}
