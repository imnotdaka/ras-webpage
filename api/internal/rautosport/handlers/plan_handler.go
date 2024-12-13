package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/internal/clients/mercadopago"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/plan"
	"github.com/mercadopago/sdk-go/pkg/preapprovalplan"
)

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

		lastID, err := r.CreatePlanDB(ctx, id, planTemp.Reason, planTemp.AutoRecurring.Frequency, planTemp.AutoRecurring.FrequencyType, planTemp.AutoRecurring.TransactionAmount)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		fmt.Println(lastID)
	}
}

func GetAllPlanHandler(r plan.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		plan, err := r.GetAllPlan(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK, plan)
	}
}
