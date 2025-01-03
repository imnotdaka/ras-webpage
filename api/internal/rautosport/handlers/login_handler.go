package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/authenticator"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/user"
)

func LoginHandler(s user.Repository, auth authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req user.LoginRequest
		err := json.NewDecoder(ctx.Request.Body).Decode(&req)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, "err valid")
		}

		email := (req.Email)
		user, err := s.GetUserByEmail(ctx, email)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		err = validatePsw(user.EncryptedPassword, req.Password)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, ErrBadRequestPassword)
			return
		}
		tokenString, err := auth.Create(&user)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		fmt.Println("Refresh_token: ", tokenString.RefreshToken)
		ctx.SetCookie("refresh_token", tokenString.RefreshToken, int(authenticator.ExpirationTimeRT), "/", "localhost", false, true)
		ctx.JSON(http.StatusOK, tokenString.AccessToken)
	}
}
