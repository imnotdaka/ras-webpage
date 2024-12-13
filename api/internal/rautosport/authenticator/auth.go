package authenticator

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/user"
)

var (
	ErrIsNotValid           = errors.New("user not valid")
	ErrInvalidSigningMethod = errors.New("unexpected signing method")
	ErrTokenIsNotValid      = errors.New("token not valid")
)

type Authenticator interface {
	CreateJWT(user *user.User) (string, error)
	VerifyJWT(tokenString string) (*jwt.Token, error)
}

type authenticator struct {
	secret string
}

func NewAuth(secret string) Authenticator {
	return &authenticator{secret: secret}
}

func (a authenticator) CreateJWT(user *user.User) (string, error) {
	claims := &jwt.MapClaims{
		"expires_at": jwt.NewNumericDate(time.Unix(1516239022, 0)),
		"user_id":    user.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.secret))
}

func (a authenticator) VerifyJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return []byte(a.secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, ErrTokenIsNotValid
	}

	return token, nil
}
