package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/session"
)

func LogOutHandler(s session.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("refresh_token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "unauth")
			return
		}
		err = s.Update(
			session.Session{
				Token:   token,
				IsValid: false,
			},
		)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK, "Logged out")
	}
}
