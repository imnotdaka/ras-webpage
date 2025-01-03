package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/internal/clients/mercadopago"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/plan"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/subscription"
	"github.com/mercadopago/sdk-go/pkg/preapproval"
)

var internalServerErr = "Internal server error"

func CreateSubscriptionHandler(c mercadopago.Client, r plan.Repository, s subscription.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var planTemp preapproval.Request
		var planRes subscription.SubscriptionRes

		err := ctx.ShouldBindJSON(&planTemp)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		res, err := c.CreateSubscription(ctx, planTemp)
		if err != nil {
			fmt.Println("uwu", err)
			ctx.JSON(http.StatusInternalServerError, internalServerErr)
			return
		}

		userID, exists := ctx.Get("user_id")
		if !exists {
			ctx.JSON(http.StatusBadRequest, "Bad Req")
			return
		}
		userIDInt := userID.(int)

		subscriptionToDB := subscription.SubscriptionToDB{
			SubscriptionID:    res.ID,
			UserID:            userIDInt,
			DateCreated:       res.DateCreated,
			NextPaymentDate:   res.NextPaymentDate,
			PreapprovalPlanID: res.PreapprovalPlanID,
			Status:            res.Status,
		}

		s.CreateSubscriptionToDB(ctx, subscriptionToDB)

		planRes.Status = res.Status
		planRes.DateCreated = res.DateCreated
		planRes.Reason = res.Reason
		planRes.TransactionAmount = res.AutoRecurring.TransactionAmount

		ctx.JSON(http.StatusOK, planRes)
	}
}

func GetSubscriptionHandler(c mercadopago.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		res, err := c.GetSubscriptionById(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		fmt.Println("getsusres", res)
		ctx.JSON(http.StatusOK, res)
	}
}

func GetSubscriptionByUserIDHandler(s subscription.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, exists := ctx.Get("user_id")
		if !exists {
			ctx.JSON(http.StatusBadRequest, "Bad Req")
			return
		}
		userIDInt := userID.(int)

		res, err := s.GetSubscriptionByUserID(ctx, userIDInt)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func UpdateMPSubscriptionHandler(c mercadopago.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req mercadopago.UpdateReq

		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if req.Status == "" {
			ctx.JSON(http.StatusBadRequest, "empty status")
			return
		}
		err = c.UpdateSubscription(ctx, req.ID, req.Status)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
	}
}

func CancelMPSubscriptionHandler(c mercadopago.Client, s subscription.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, exists := ctx.Get("user_id")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, "unauth")
			return
		}
		sub, err := s.GetSubscriptionByUserID(ctx, id.(int))
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		err = c.UpdateSubscription(ctx, sub.ID, mercadopago.Cancelled)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK, "Subscription cancelled")
	}
}
