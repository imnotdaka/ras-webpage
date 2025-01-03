package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
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

type refreshRes struct {
	AccessToken string `json:"access_token"`
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
