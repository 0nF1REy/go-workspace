package middleware

import (
	"net/http"
	"strings"

	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/model"
	"github.com/gin-gonic/gin"
)

type TokenValidator interface {
	ValidateToken(tokenValue string) (*model.AuthClaims, error)
}

func JWTAuthMiddleware(validator TokenValidator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := strings.TrimSpace(ctx.GetHeader("Authorization"))
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token de acesso ausente"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || strings.TrimSpace(parts[1]) == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "formato de autorização inválido"})
			return
		}

		claims, err := validator.ValidateToken(parts[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido ou expirado"})
			return
		}

		ctx.Set("jwt_claims", claims)
		ctx.Next()
	}
}
