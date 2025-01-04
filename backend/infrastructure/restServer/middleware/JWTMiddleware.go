package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/revazcus/task-tracker/backend/infrastructure/errors"
	loggerInterface "github.com/revazcus/task-tracker/backend/infrastructure/logger/interface"
	jwtServiceInterface "github.com/revazcus/task-tracker/backend/infrastructure/security/jwtService/interface"
	"net/http"
	"strings"
)

type JWTMiddleware struct {
	logger     loggerInterface.Logger
	jwtService jwtServiceInterface.JWTService
}

func NewJWTMiddleware(logger loggerInterface.Logger, jwtService jwtServiceInterface.JWTService) *JWTMiddleware {
	return &JWTMiddleware{
		logger:     logger,
		jwtService: jwtService,
	}
}

func (r *JWTMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := r.extractAuthToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: " + err.Error(),
			})
			return
		}

		if !r.jwtService.Verify(token) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: Invalid token",
			})
			return
		}

		reqCtx := ctx.Request.Context()

		userContext, err := r.jwtService.FillCtxWithParams(reqCtx, token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: Failed to fill context",
			})
			return
		}

		ctx.Request = ctx.Request.WithContext(userContext)

		ctx.Next()
	}
}

func (r *JWTMiddleware) extractAuthToken(ctx *gin.Context) (string, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return "", errors.NewError("SYS", "Unauthorized")
	}

	if strings.HasPrefix(token, "Bearer ") {
		token = token[7:]
	}

	return token, nil
}
