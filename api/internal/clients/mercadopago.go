package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/plan"
)

type AutoRecurring struct {
	Frequency              int    `json:"frequency"`
	FrequencyType          string `json:"frequency_type"`
	BillingDay             int    `json:"billing_day"`
	BillingDayProportional bool   `json:"billing_day_proportional"`
	TransactionAmount      int    `json:"transaction_amount"`
	CurrencyID             string `json:"currency_id"`
}

type PaymentMethodsAllowed struct {
	Id string `json:"id"`
}

type PreApprovalPlan struct {
	ID                    string                `json:"id"`
	Reason                string                `json:"reason"`
	AutoRecurring         AutoRecurring         `json:"auto_recurring"`
	PaymentMethodsAllowed PaymentMethodsAllowed `json:"payment_methods_allowed"`
	BackURL               string                `json:"back_url"`
}

func CreatePlan(r plan.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := os.Getenv("ACCESS_TOKEN")
		var planTemp PreApprovalPlan

		err := ctx.ShouldBindJSON(&planTemp)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		url := "https://api.mercadopago.com/preapproval_plan"

		jsonData, err := json.Marshal(planTemp)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}

		req.Header.Set("Authorization", "Bearer "+accessToken)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		fmt.Println("res:", res.Body)

		defer res.Body.Close()

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		if res.StatusCode != http.StatusCreated {
			ctx.JSON(res.StatusCode, err)
			return
		}

		var parsedPlan PreApprovalPlan

		// var apiResponse map[string]interface{}
		err = json.Unmarshal(resBody, &parsedPlan)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		lastID, err := plan.Repository.CreatePlanDB(r, &parsedPlan.ID, &parsedPlan.Reason, &parsedPlan.AutoRecurring.Frequency, &parsedPlan.AutoRecurring.FrequencyType, &parsedPlan.AutoRecurring.TransactionAmount)

		fmt.Printf("%v+ \n", &parsedPlan)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		fmt.Println(lastID)

		ctx.JSON(http.StatusOK, "Pre-approval plan created successfully")
	}
}

func GetAll(r plan.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		plan, err := plan.Repository.GetAll(r)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		fmt.Printf("%v+ \n", plan)
		ctx.JSON(http.StatusOK, plan)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, "order created")
	}
}

func Failure() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, "order created")
	}
}

func Pending() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, "order created")
	}
}

func Success() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Success")
	}
}

func Webhook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Webhook")
	}
}
