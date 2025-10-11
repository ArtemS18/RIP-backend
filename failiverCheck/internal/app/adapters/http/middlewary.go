package http

import (
	"failiverCheck/internal/app/dto"
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

func (h *Handler) AuthoMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
		login, ok := claims["login"].(string)
		if !ok {
			h.errorHandler(ctx, http.StatusUnauthorized, fmt.Errorf("bad jwt credentials"))
			return
		}
		ctx.Set("userDTO", dto.UserDTO{ID: uint(userId), IsModerator: is_moderator, Login: login})
		ctx.Next()
	}
}

func (h *Handler) RoleValidateMiddleware(allowedRoles ...Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := h.GetUserDTO(ctx)
		if err != nil {
			h.errorHandler(ctx, 404, err)
			return
		}
		var role Role
		if user.IsModerator {
			role = ModeratorRole
		}
		if !slices.Contains(allowedRoles, ModeratorRole) {
			allowedRoles = append(allowedRoles, ModeratorRole)
		}

		if !slices.Contains(allowedRoles, role) {
			h.errorHandler(ctx, http.StatusForbidden, fmt.Errorf("not allowed role"))
			return
		}
		ctx.Next()
	}
}
func (h *Handler) GetUserID(ctx *gin.Context) uint {
	dto, err := h.GetUserDTO(ctx)
	if err != nil {
		h.errorHandler(ctx, 404, err)
		return 0
	}
	return uint(dto.ID)
}

func (h *Handler) GetUserDTO(ctx *gin.Context) (dto.UserDTO, error) {
	userRaw, ok := ctx.Get("userDTO")
	if !ok {
		return dto.UserDTO{}, fmt.Errorf("user not found")
	}
	user, ok := userRaw.(dto.UserDTO)
	if !ok {
		return dto.UserDTO{}, fmt.Errorf("invalid user id type: expected int, got %T", user)
	}
	return user, nil
}
