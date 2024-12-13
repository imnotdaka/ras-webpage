package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/authenticator"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/user"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrBadRequestCreateUser = errors.New("todos los campos son requeridos")
	ErrBadRequestParseID    = errors.New("el ID debe ser un numero")
	ErrBadRequestEmail      = errors.New("rellene el campo de email")
	ErrBadRequestPassword   = errors.New("el email o la contraseña es incorrecta")
)

var (
	validateEmail = regexp.MustCompile(`^([\w-.])+@([\w-])+.+([\w-.]{2,6})$`)
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

		tokenString, err := auth.CreateJWT(user)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		fmt.Println("JTW token: ", tokenString)
		ctx.SetCookie("x-jwt-token", tokenString, 3600, "/", "localhost", false, false)
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
			ctx.JSON(http.StatusBadRequest, ErrBadRequestParseID)
			return
		}
		user, err := r.GetUserById(ctx, id)
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
		deletedID, err := r.DeleteUser(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, deletedID)
	}
}

func Login(s user.Repository, auth authenticator.Authenticator) gin.HandlerFunc {
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
		tokenString, err := auth.CreateJWT(&user)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		fmt.Println("JTW token: ", tokenString)
		ctx.SetCookie("x-jwt-token", tokenString, 3600, "/", "localhost", false, false)
		ctx.JSON(http.StatusOK, "OK")
	}
}

func validateUser(req user.RegisterReq) error {
	if req.FirstName == "" ||
		req.LastName == "" ||
		req.Email == "" ||
		req.Password == "" {
		return ErrBadRequestCreateUser
	}
	if !validateEmail.MatchString(req.Email) {
		return fmt.Errorf("email invalido")
	}
	return nil
}

func validatePsw(encpw string, reqpsw string) error {
	fmt.Println("encpw:", encpw, "reqpsw:", reqpsw)
	err := bcrypt.CompareHashAndPassword([]byte(encpw), []byte(reqpsw))
	if err != nil {
		return err
	}
	return nil
}
