package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	loggerInterface "task-tracker/infrastructure/logger/interface"
	jwtServiceInterface "task-tracker/infrastructure/security/jwtService/interface"
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

		ctx.Set("userContext", userContext)

		ctx.Next()
	}
}

func (r *JWTMiddleware) extractAuthToken(ctx *gin.Context) (string, error) {
	token := ctx.GetString("Authorization")
	if token == "" {
		return "", errors.New("unauthorized")
	}

	if strings.HasPrefix(token, "Bearer ") {
		token = token[7:]
	}

	return token, nil
}
