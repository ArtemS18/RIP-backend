package handler

import (
	"failiverCheck/internal/app/schemas"
	"fmt"
	"net/http"
	"slices"
	"strings"

	utils "failiverCheck/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Role string

var (
	ModeratorRole Role = "MODERATOR"
	UserRole      Role = "USER"
)

func (h *Handler) AuthMiddleware(allowedRoles ...Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		existRole, ok := ctx.Get("userRole")
		if ok {
			roleStr, _ := existRole.(string)
			role := Role(roleStr)
			if !slices.Contains(allowedRoles, ModeratorRole) {
				allowedRoles = append(allowedRoles, ModeratorRole)
			}
			if !slices.Contains(allowedRoles, role) {
				h.errorHandler(ctx, http.StatusForbidden, fmt.Errorf("not allowed role"))
				return
			}
			ctx.Next()
		}
		jwtPrefix := "Bearer "
		jwtStr := ctx.Request.Header.Get("Authorization")
		if !strings.HasPrefix(jwtStr, jwtPrefix) {
			ctx.JSON(http.StatusUnauthorized, schemas.Error{
				Status:      "error",
				Description: "access token not found",
			})
			ctx.Abort()
			return
		}
		jwtTokenStr := jwtStr[len(jwtPrefix):]
		claims, err := utils.ValidateJwtToken(jwtTokenStr, h.Config.JWT.SecretKey)
		if err != nil {
			h.errorHandler(ctx, http.StatusUnauthorized, err)
			return
		}
		userId, ok := claims["sub"].(float64)
		logrus.Info(userId)
		if !ok {
			h.errorHandler(ctx, http.StatusUnauthorized, fmt.Errorf("bad jwt credentials"))
			return
		}
		is_moderator, ok := claims["is_moderator"].(bool)
		logrus.Info(is_moderator)
		if !ok {
			h.errorHandler(ctx, http.StatusUnauthorized, fmt.Errorf("bad jwt credentials"))
			return
		}
		role := UserRole
		if is_moderator {
			role = ModeratorRole
		}
		if !slices.Contains(allowedRoles, ModeratorRole) {
			allowedRoles = append(allowedRoles, ModeratorRole)
		}

		if !slices.Contains(allowedRoles, role) {
			h.errorHandler(ctx, http.StatusForbidden, fmt.Errorf("not allowed role"))
			return
		}
		ctx.Set("userId", userId)
		ctx.Set("userRole", role)
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
	id, ok := raw.(float64)
	if !ok {
		h.errorHandler(ctx, 404, fmt.Errorf("invalid user id type: expected int, got %T", raw)) // Используем %T для получения типа
		return 0
	}
	return uint(id)
}
