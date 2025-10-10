package handler

import (
	"failiverCheck/internal/app/schemas"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) errorHandler(ctx *gin.Context, errorCode int, err error) {
	log.Error(err.Error())
	ctx.JSON(errorCode, schemas.Error{
		Status:      "error",
		Description: err.Error(),
	})
	ctx.Abort()

}

func (h *Handler) successHandler(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, schemas.OKResponse{
		Status: "ok",
		Data:   data,
	})

}

func (h *Handler) getIntParam(ctx *gin.Context, param string) int {
	raw := ctx.Param(param)
	paramInt, err := strconv.Atoi(raw)
	if err != nil {
		h.errorHandler(ctx, 400, fmt.Errorf("invaid param id=%s", raw))
		return 0
	}
	return paramInt

}

func (h *Handler) getFileHeaders(ctx *gin.Context) (string, int64) {
	contentType := ctx.Request.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	contentLengthStr := ctx.Request.Header.Get("Content-Length")
	if contentLengthStr == "" {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("content-Length header is required"))
		return "", 0
	}
	fileSize, err := strconv.ParseInt(contentLengthStr, 10, 64)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("invalid Content-Length header"))
		return "", 0
	}
	return contentType, fileSize
}

func (h *Handler) validateFields(ctx *gin.Context, obj any) {
	if err := ctx.BindJSON(obj); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		ctx.Abort()

	}
	validate := validator.New()
	if err := validate.Struct(obj); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		ctx.Abort()
	}
}
