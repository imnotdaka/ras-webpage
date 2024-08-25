package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/user"
)

func CreateUserHandler(r user.Repository) gin.HandlerFunc {

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
		tokenString, err := createJWT(&user)
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
		id := ctx.Param("id")
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

func createJWT(user *user.User) (string, error) {
	claims := &jwt.MapClaims{
		"ExpiresAt": jwt.NewNumericDate(time.Unix(1516239022, 0)),
		"userID":    user.ID,
	}
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("JWT Created")
	return token.SignedString([]byte(secret))
}

func WithJWTAuth(handlerfunc gin.HandlerFunc, s user.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Printf("JWT Auth")
		tokenString := ctx.GetHeader("x-jwt-token")
		token, err := user.JWTValidation(tokenString)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "err validation")
			return
		}
		if !token.Valid {
			ctx.JSON(http.StatusInternalServerError, "!token.valid")
			return
		}

		userID := ctx.Param("id")
		user, err := s.GetUserById(userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if user.ID != claims["userID"] {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, "OK")
	}
}
