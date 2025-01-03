package mercadopago

import "time"

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

type UpdateReq struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type Data struct {
	ID string `json:"id"`
}

type RequestBodyWebhook struct {
	Action        string    `json:"action"`
	ApplicationID int       `json:"application_id"`
	Data          Data      `json:"data"`
	Date          time.Time `json:"date"`
	Entity        string    `json:"entity"`
	ID            int       `json:"id"`
	Type          string    `json:"type"`
}

var (
	SubscriptionPreapproval = "subscription_preapproval"
	PlanPreapproval         = "subscription_preapproval_plan"
	Updated                 = "updated"
	Cancelled               = "cancelled"
)
