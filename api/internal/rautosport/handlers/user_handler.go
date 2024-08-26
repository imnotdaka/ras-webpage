package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/authenticator"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/user"
)

func CreateUserHandler(r user.Repository, auth authenticator.Authenticator) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var user user.User
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		res, err := r.CreateUser(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		user.ID = int(res)

		tokenString, err := auth.CreateJWT(&user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		fmt.Println("JTW token: ", tokenString)

		ctx.JSON(http.StatusOK, res)
	}

}

func GetUserHandler() gin.HandlerFunc {

	return func(ctx *gin.Context) {}
}

func GetUserByIdHandler(r user.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		user, err := r.GetUserById(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func UpdateUserHandler(r user.Repository) gin.HandlerFunc {

	return func(ctx *gin.Context) {}
}

func DeleteUserHandler(r user.Repository) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		deletedID, err := r.DeleteUser(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, deletedID)
	}
}

func Login(s user.Repository, auth authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("x-jwt-token")
		userID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "err validation")
			return
		}

		if tokenString != "" {
			err := auth.VerifyJWT(tokenString, userID)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, "err validation")
				return
			}
		}

		user, err := s.GetUserById(userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}
