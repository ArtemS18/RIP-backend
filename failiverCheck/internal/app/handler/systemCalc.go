package handler

import (
	"failiverCheck/internal/app/dto"
	"failiverCheck/internal/app/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetSystemCalc(ctx *gin.Context) {
	id := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	systemCalc, err := h.Repository.GetSystemCalcById(uint(id))
	sysCalcsResp := dto.ToSystemCalculationDTO(systemCalc)
	if err != nil {
		h.errorHandler(ctx, 404, err)
		return
	}

	h.successHandler(ctx, 200, sysCalcsResp)
}
func (h *Handler) GetSystemCalcList(ctx *gin.Context) {
	var filters dto.SystemCalcFilters
	if errQ := ctx.BindQuery(&filters); errQ != nil {
		h.errorHandler(ctx, http.StatusBadRequest, errQ)
		return
	}
	sysCalcs, err := h.Repository.GetSystemCalcList(filters)
	sysCalcsResp := dto.ToSystemCalculationListDTO(sysCalcs)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}
	h.successHandler(ctx, http.StatusOK, sysCalcsResp)
}

func (h *Handler) GetSystemCalcBucket(ctx *gin.Context) {
	var userId uint = h.GetUserID(ctx)
	if ctx.IsAborted() {
		return
	}
	bucket, err := h.Repository.GetCurrentSysCalcAndCount(userId)
	if err != nil {
		h.errorHandler(ctx, 400, err)
		return
	}
	h.successHandler(ctx, http.StatusOK, bucket)
}

func (h *Handler) UpdateSystemCalc(ctx *gin.Context) {
	sysCalcId := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	var data schemas.UpdateSystemCalcFields
	h.validateFields(ctx, &data)
	if ctx.IsAborted() {
		return
	}
	system, err := h.Repository.UpdateSystemCalc(uint(sysCalcId), dto.UpdateSystemCalcDTO{SystemName: data.SystemName})
	systemDto := dto.ToSystemCalculationDTO(system)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	h.successHandler(ctx, 200, systemDto)
}

func (h *Handler) UpdateSystemCalcStatusToFormed(ctx *gin.Context) {
	sysCalcId := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	system, err := h.Repository.UpdateSystemCalcStatusToFormed(uint(sysCalcId))
	dto := dto.ToSystemCalculationDTO(system)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	h.successHandler(ctx, 200, dto)
}

func (h *Handler) UpdateSystemCalcStatusModerator(ctx *gin.Context) {
	moderatorId := uint(1)
	sysCalcId := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	var data schemas.UpdateSystemCalcStatus
	h.validateFields(ctx, &data)
	if ctx.IsAborted() {
		return
	}
	system, err := h.Repository.UpdateSystemCalcStatusModerator(uint(sysCalcId), moderatorId, data.Command)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	dto := dto.ToSystemCalculationDTO(system)
	h.successHandler(ctx, 200, dto)
}

func (h *Handler) DeleteSystemCalc(ctx *gin.Context) {
	sysCalcId := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	err := h.Repository.DeleteSystemCalc(uint(sysCalcId))
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	h.successHandler(ctx, 204, nil)
}
