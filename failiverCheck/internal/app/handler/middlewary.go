package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.Set("userId", 1)
		ctx.Next()
	}
}

func (h *Handler) GetUserID(ctx *gin.Context) uint {
	raw, ok := ctx.Get("userId")
	if !ok {
		h.errorHandler(ctx, 400, fmt.Errorf("user id not found"))
		return 0
	}
	logrus.Info(raw)
	id, ok := raw.(int)
	if !ok {
		h.errorHandler(ctx, ctx.Writer.Status(), fmt.Errorf("invalid user id type: expected int, got %T", raw)) // Используем %T для получения типа
		return 0
	}
	return uint(id)
}
