package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/cmd/config"
	"github.com/imnotdaka/RAS-webpage/internal/clients/mercadopago"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/plan"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/subscription"
)

func WebhookHandler(c mercadopago.Client, s subscription.Repository, p plan.Repository, cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		err := SecretKeyValidate(ctx, cfg)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			return
		}
		var req mercadopago.RequestBodyWebhook

		err = ctx.ShouldBindJSON(&req)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		fmt.Println("req", req)
		if req.Type == mercadopago.SubscriptionPreapproval {
			fmt.Println("processing subscription")

			switch req.Action {
			case mercadopago.Updated:
				res, err := c.GetSubscriptionById(ctx, req.Data.ID)
				if err != nil {
					fmt.Println(err)
					ctx.JSON(http.StatusInternalServerError, err)
					return
				}
				err = s.UpdateSubscription(ctx, subscription.UpdateReq{
					NextPaymentDate: res.NextPaymentDate,
					Status:          res.Status,
					SubscriptionID:  res.ID,
				})
				if err != nil {
					fmt.Println(err)
					ctx.JSON(http.StatusInternalServerError, err)
					return
				}
			}
		}
		if req.Type == mercadopago.PlanPreapproval {
			fmt.Println("processing preapprovalplan")
			switch req.Action {
			case mercadopago.Updated:
				res, err := c.GetPlan(ctx, req.Data.ID)
				if err != nil {
					fmt.Println(err)
					ctx.JSON(http.StatusInternalServerError, err)
					return
				}
				err = p.UpdatePlan(ctx, &plan.PreApprovalPlan{
					ID:     res.ID,
					Reason: res.Reason,
					AutoRecurring: plan.AutoRecurring{
						TransactionAmount: res.AutoRecurring.TransactionAmount,
					},
				})
				if err != nil {
					fmt.Println(err)
					ctx.JSON(http.StatusInternalServerError, err)
					return
				}
			}

		}
	}
}

func SecretKeyValidate(ctx *gin.Context, cfg *config.Config) error {
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
	secret := cfg.SecretKey

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
