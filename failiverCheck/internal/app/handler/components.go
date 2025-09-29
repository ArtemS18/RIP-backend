package handler

import (
	"failiverCheck/internal/app/ds"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TemplateData struct {
	Components   []ds.Component
	CountQuery   int
	CurrentCount int
	SearchQuery  string
	SystemCalcId int
}

func (h *Handler) GetComponent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.errorHandler(ctx, 400, err)
	}
	log.Info(id)
	order, err := h.Repository.GetComponentById(id)
	if err != nil {
		h.errorHandler(ctx, 500, err)
	}

	ctx.HTML(http.StatusOK, "component.html", gin.H{
		"component": order,
	})
}

func (h *Handler) GetComponents(ctx *gin.Context) {
	var orders []ds.Component
	var err error
	var searchQuery string
	param, ok := ctx.Get("search")

	if !ok {
		searchQuery = ctx.Query("search")
	} else {
		searchQuery, ok = param.(string)
		if !ok {
			h.errorHandler(ctx, 400, fmt.Errorf("bad param"))
		}
	}
	count, err := h.Repository.GetCountComnponents(1)
	if err != nil {
		h.errorHandler(ctx, 500, err)
	}

	log.Info(searchQuery)
	if searchQuery == "" {
		orders, err = h.Repository.GetComponents()
		if err != nil {
			h.errorHandler(ctx, 400, err)
		}
	} else {
		orders, err = h.Repository.GetComponentsByTitle(searchQuery)
		if err != nil {
			h.errorHandler(ctx, 500, err)
		}
	}
	systemCalc, err := h.Repository.GetSystemCalc(1)
	var systemCalcId int
	if err != nil {
		log.Error(err)
		systemCalcId = 0
	} else {
		systemCalcId = int(systemCalc.ID)
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"data": TemplateData{
			Components:   orders,
			SearchQuery:  searchQuery,
			SystemCalcId: systemCalcId,
			CurrentCount: count},
	})
}
