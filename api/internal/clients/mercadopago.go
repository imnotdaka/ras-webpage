package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AutoRecurring struct {
	Frequency              int    `json:"frequency"`
	FrequencyType          string `json:"frequency_type"`
	BillingDay             int    `json:"billing_day"`
	BillingDayProportional bool   `json:"billing_day_proportional"`
	TransactionAmout       int    `json:"transaction_amount"`
	CurrencyID             string `json:"currency_id"`
}

type PaymentMethodsAllowed struct {
	Id string `json:"id"`
}

type PreApprovalPlan struct {
	Reason                string                `json:"reason"`
	AutoRecurring         AutoRecurring         `json:"auto_recurring"`
	PaymentMethodsAllowed PaymentMethodsAllowed `json:"payment_methods_allowed"`
	BackURL               string                `json:"back_url"`
}

func CreatePlan() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := os.Getenv("ACCESS_TOKEN")
		var plan PreApprovalPlan

		err := ctx.ShouldBindJSON(&plan)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		url := "https://api.mercadopago.com/preapproval_plan"

		jsonData, err := json.Marshal(plan)
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
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}

		if res.StatusCode != http.StatusCreated {
			ctx.JSON(res.StatusCode, err)
			return
		}

		var apiResponse map[string]interface{}
		err = json.Unmarshal(resBody, &apiResponse)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		fmt.Println("parsed res: ", &apiResponse)

		ctx.JSON(http.StatusOK, "Pre-approval plan created successfully")
	}
}

func GetPlan() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := os.Getenv("ACCESS_TOKEN")
		url := "https://api.mercadopago.com/preapproval_plan/search"

		req, err := http.NewRequest("GET", url, nil)
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

		defer res.Body.Close()

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		var apiResponse map[string]interface{}
		err = json.Unmarshal(resBody, &apiResponse)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		fmt.Println("parsed res: ", &apiResponse)
		ctx.JSON(http.StatusOK, &apiResponse)
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
