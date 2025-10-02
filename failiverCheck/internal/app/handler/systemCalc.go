package handler

import (
	"failiverCheck/internal/app/dto"
	"failiverCheck/internal/app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetSystemCalc(ctx *gin.Context) {
	id := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	systemCalc, err := h.Repository.GetSystemCalcById(uint(id))
	if err != nil {
		h.errorHandler(ctx, 404, err)
		return
	}

	h.successHandler(ctx, 200, systemCalc)
}
func (h *Handler) GetSystemCalcList(ctx *gin.Context) {
	var filters dto.SystemCalcFilters
	if errQ := ctx.BindQuery(&filters); errQ != nil {
		h.errorHandler(ctx, http.StatusBadRequest, errQ)
		return
	}
	sysCalcs, err := h.Repository.GetSystemCalcList(filters)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}
	h.successHandler(ctx, http.StatusOK, sysCalcs)
}
func (h *Handler) DeleteComponentFromSystemCalc(ctx *gin.Context) {
	var err error
	sysCalcId := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	strId := ctx.PostForm("component_id")
	componentId, err := strconv.Atoi(strId)
	if err != nil {
		h.errorHandler(ctx, 400, err)
		return
	}

	err = h.Repository.DeleteComponentFromSystemCalc(uint(sysCalcId), uint(componentId))
	if err != nil {
		h.errorHandler(ctx, 404, err)
		return
	}
	h.successHandler(ctx, 204, nil)
}

func (h *Handler) DeleteSystemCalc(ctx *gin.Context) {
	idStr := ctx.PostForm("system_id")
	sysCalcId, err := strconv.Atoi(idStr)
	if err != nil {
		h.errorHandler(ctx, 400, err)
		return
	}
	logrus.Info(idStr)
	err = h.Repository.DeleteSystemCalc(uint(sysCalcId))
	if err != nil {
		h.errorHandler(ctx, 400, err)
		return
	}
	h.successHandler(ctx, 204, nil)
}

func (h *Handler) GetSystemCalcBucket(ctx *gin.Context) {
	userId := uint(1)
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
	var data models.UpdateSystemCalcFields
	if err := ctx.ShouldBindJSON(&data); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	err := h.Repository.UpdateSystemCalc(uint(sysCalcId), dto.UpdateSystemCalcDTO{SystemName: data.SystemName})
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	h.successHandler(ctx, 204, nil)
}

func (h *Handler) UpdateSystemCalcStatusToFormed(ctx *gin.Context) {
	sysCalcId := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	err := h.Repository.UpdateSystemCalcStatusToFormed(uint(sysCalcId))
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	h.successHandler(ctx, 204, nil)
}

func (h *Handler) UpdateSystemCalcStatusModerator(ctx *gin.Context) {
	moderatorId := uint(1)
	sysCalcId := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	var data models.UpdateSystemCalcStatus
	if err := ctx.ShouldBindJSON(&data); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	err := h.Repository.UpdateSystemCalcStatusModerator(uint(sysCalcId), moderatorId, data.Command)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	h.successHandler(ctx, 204, nil)
}
