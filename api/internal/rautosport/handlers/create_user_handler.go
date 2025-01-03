package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/authenticator"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/user"
)

func CreateUserHandler(r user.Repository, auth authenticator.Authenticator) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req user.RegisterReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			slog.Error("bad request error", "error", err)
			ctx.JSON(http.StatusBadRequest, ErrBadRequestCreateUser)
			return
		}
		fmt.Println(req)
		err = validateUser(req)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		user, err := user.NewAccount(req.FirstName, req.LastName, req.Email, req.Password)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		res, err := r.CreateUser(ctx, user)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		user.ID = int(res)

		tokenString, err := auth.Create(user)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		fmt.Println("RefreshToken: ", tokenString.RefreshToken)
		ctx.SetCookie("refresh_token", tokenString.RefreshToken, int(authenticator.ExpirationTimeRT), "/", "localhost", false, false)
		ctx.JSON(http.StatusOK, tokenString.AccessToken)
	}
}
