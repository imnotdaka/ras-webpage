package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/internal/clients/mercadopago"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/plan"
	"github.com/mercadopago/sdk-go/pkg/preapproval"
	"github.com/mercadopago/sdk-go/pkg/preapprovalplan"
)

var internalServerErr = "Internal server error"

func CreatePlanHandler(c mercadopago.Client, r plan.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var planTemp preapprovalplan.Request

		err := ctx.ShouldBindJSON(&planTemp)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		id, err := c.CreatePlan(ctx, planTemp)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, internalServerErr)
			return
		}

		lastID, err := r.CreatePlanDB(id, planTemp.Reason, planTemp.AutoRecurring.Frequency, planTemp.AutoRecurring.FrequencyType, planTemp.AutoRecurring.TransactionAmount)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		fmt.Println(lastID)
	}
}

func CreateSuscriptionHandler(c mercadopago.Client, r plan.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var planTemp preapproval.Request

		err := ctx.ShouldBindJSON(&planTemp)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		_, err = c.CreateSuscription(ctx, planTemp)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, internalServerErr)
			return
		}

		// lastID, err := r.CreatePlanDB(id, planTemp.Reason, planTemp.AutoRecurring.Frequency, planTemp.AutoRecurring.FrequencyType, planTemp.AutoRecurring.TransactionAmount)
		// if err != nil {
		// 	ctx.JSON(http.StatusBadRequest, err)
		// 	return
		// }

		// fmt.Println(lastID)

	}
}
