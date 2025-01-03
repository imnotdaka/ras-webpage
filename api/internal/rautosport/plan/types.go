package plan

type AutoRecurring struct {
	Frequency              int     `json:"frequency"`
	FrequencyType          string  `json:"frequency_type"`
	BillingDay             int     `json:"billing_day"`
	BillingDayProportional bool    `json:"billing_day_proportional"`
	TransactionAmount      float64 `json:"transaction_amount"`
	CurrencyID             string  `json:"currency_id"`
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
