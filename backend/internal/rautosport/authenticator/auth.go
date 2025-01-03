package authenticator

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/session"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/user"
)

var (
	ExpirationTimeAT = time.Minute * 15
	ExpirationTimeRT = time.Hour * 1
)

var (
	ErrIsNotValid           = errors.New("user not valid")
	ErrInvalidSigningMethod = errors.New("unexpected signing method")
	ErrTokenIsNotValid      = errors.New("token not valid")
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Authenticator interface {
	Create(user *user.User) (Tokens, error)
	Verify(tokenString string) (*jwt.Token, error)
	Refresh(ctx context.Context, token string) (string, error)
}

type authenticator struct {
	session  session.Repository
	atSecret string
	rtSecret string
}

func NewAuth(atSecret, rtSecret string, repo session.Repository) *authenticator {
	return &authenticator{
		atSecret: atSecret,
		rtSecret: rtSecret,
		session:  repo,
	}
}

func (a authenticator) Create(user *user.User) (Tokens, error) {
	atClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Unix() + int64(ExpirationTimeAT.Seconds()),
	})
	AccessToken, err := atClaims.SignedString([]byte(a.atSecret))
	if err != nil {
		return Tokens{}, err
	}

	rtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Unix() + int64(ExpirationTimeRT.Seconds()),
		"iat":     time.Now().Unix(),
	})

	refreshToken, err := rtClaims.SignedString([]byte(a.rtSecret))
	if err != nil {
		return Tokens{}, err
	}

	err = a.session.Create(session.Session{
		Token:  refreshToken,
		UserID: user.ID,
	})
	if err != nil {
		slog.Error("session create error", "error", err)
		return Tokens{}, err
	}

	return Tokens{
		AccessToken:  AccessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a authenticator) Verify(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return []byte(a.atSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, ErrTokenIsNotValid
	}
	return token, nil
}

func (a authenticator) Refresh(ctx context.Context, token string) (string, error) {
	s, err := a.session.Get(ctx, token)
	if err != nil {
		slog.Error("Session get in refresh", "error", err)
		return "", err
	}
	if !s.IsValid {
		slog.Error("Session not valid")
		return "", ErrIsNotValid
	}

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			slog.Error("Error parsing JWT")
			return nil, http.ErrAbortHandler
		}
		return []byte(a.rtSecret), nil
	})

	if err != nil || !t.Valid {
		return "", errors.New("Unauthorized")
	}
	claims := t.Claims.(jwt.MapClaims)

	atClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": claims["user_id"],
		"exp":     time.Now().Unix() + int64(ExpirationTimeAT.Seconds()),
	})

	accessToken, err := atClaims.SignedString([]byte(a.atSecret))
	if err != nil {
		slog.Error("signedstring", "error", err)
		return "", err
	}

	return accessToken, nil
}
