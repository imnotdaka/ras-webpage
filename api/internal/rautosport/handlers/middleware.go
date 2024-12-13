package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/authenticator"
)

func JWTMiddleware(auth authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("x-jwt-token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}
		token, err := auth.VerifyJWT(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := int(claims["user_id"].(float64))

		ctx.Set("user_id", userID)
	}
}
