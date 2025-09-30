package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetSystemCalc(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.errorHandler(ctx, 400, err)
		return
	}
	systemCalc, err := h.Repository.GetSystemCalcById(uint(id))
	if err != nil {
		h.errorHandler(ctx, 404, err)
		return
	}
	componentsInCalc, err := h.Repository.GetComponentsInSystemCalc(uint(systemCalc.ID))
	if err != nil {
		h.errorHandler(ctx, 500, err)
		return
	}

	ctx.HTML(http.StatusOK, "system_calc.html", gin.H{
		"components": componentsInCalc,
		"systemCalc": systemCalc,
	})
}

func (h *Handler) AddComponentInSystemCalc(ctx *gin.Context) {
	var err error
	strId := ctx.PostForm("component_id")
	componentId, err := strconv.Atoi(strId)
	if err != nil {
		h.errorHandler(ctx, 400, err)
		return
	}
	err = h.Repository.AddComponentInSystemCalc(uint(componentId), 1)
	if err != nil {
		h.errorHandler(ctx, 500, err)
		return
	}
	ctx.Redirect(301, "/components?componentName="+ctx.PostForm("componentName"))
}

func (h *Handler) DeleteComponentFromSystemCalc(ctx *gin.Context) {
	var err error
	idStr := ctx.Param("id")
	sysCalcId, err := strconv.Atoi(idStr)
	if err != nil {
		h.errorHandler(ctx, 400, err)
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
	ctx.Redirect(301, "/components")
}

func (h *Handler) DeleteSystemCalc(ctx *gin.Context) {
	idStr := ctx.PostForm("system_id")
	sysCalcId, err := strconv.Atoi(idStr)
	if err != nil {
		h.errorHandler(ctx, 400, err)
	}
	logrus.Warn(idStr)
	err = h.Repository.DeleteSystemCalc(uint(sysCalcId))
	if err != nil {
		h.errorHandler(ctx, 400, err)
	}
	ctx.Redirect(301, "/components")

}
