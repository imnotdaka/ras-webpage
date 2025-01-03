package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/authenticator"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/user"
)

type authMeRes struct {
	User  user.User `json:"user"`
	Token string    `json:"token"`
}

func RefreshHandler(a authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("refresh_token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "unauth")
			return
		}

		accessToken, err := a.Refresh(ctx, token)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, "internal server error")
			return
		}

		ctx.JSON(http.StatusOK, refreshRes{
			AccessToken: accessToken,
		})
	}
}

func AuthMeHandler(u user.Repository, a authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("refresh_token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "unauth")
			return
		}

		signedToken, err := a.Refresh(ctx, token)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusUnauthorized, "unauth")
			return
		}

		accessToken, err := a.Verify(signedToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "unauth")
			return
		}

		user, err := u.GetUserById(ctx, int(accessToken.Claims.(jwt.MapClaims)["user_id"].(float64)))
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, "internal server error")
			return
		}

		fmt.Println(user)
		ctx.JSON(http.StatusOK, authMeRes{
			User:  user,
			Token: signedToken,
		})
	}
}
