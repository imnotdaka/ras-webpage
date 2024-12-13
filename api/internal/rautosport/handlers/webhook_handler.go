package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/internal/clients/mercadopago"
)

func Webhook(c mercadopago.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		err := SecretKeyValidate(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			return
		}
		var req struct {
			Data map[string]string `json:"data"`
		}

		err = ctx.ShouldBindJSON(&req)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		fmt.Printf("%+v\n", req)
		res, _ := c.GetSubscriptionById(ctx, req.Data["id"])
		fmt.Printf("%+v", res)
	}
}

func SecretKeyValidate(ctx *gin.Context) error {
	xSignature := ctx.GetHeader("x-signature")
	xRequestId := ctx.GetHeader("x-request-id")

	queryParams := ctx.Request.URL.Query()

	dataID := queryParams.Get("data.id")

	parts := strings.Split(xSignature, ",")

	var ts, hash string

	for _, part := range parts {
		keyValue := strings.SplitN(part, "=", 2)
		if len(keyValue) == 2 {
			key := strings.TrimSpace(keyValue[0])
			value := strings.TrimSpace(keyValue[1])
			if key == "ts" {
				ts = value
			} else if key == "v1" {
				hash = value
			}
		}
	}
	secret := os.Getenv("SECRET")

	manifest := fmt.Sprintf("id:%v;request-id:%v;ts:%v;", dataID, xRequestId, ts)

	hmac := hmac.New(sha256.New, []byte(secret))
	hmac.Write([]byte(manifest))

	sha := hex.EncodeToString(hmac.Sum(nil))

	if sha != hash {
		fmt.Println("HMAC verification failed")
		return errors.New("HMAC verification failed")
	}
	fmt.Println("HMAC verification successfull")
	return nil
}
