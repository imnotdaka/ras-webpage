package subscription

import "time"

type SubscriptionRes struct {
	Status            string    `json:"status"`
	DateCreated       time.Time `json:"date_created"`
	Reason            string    `json:"reason"`
	TransactionAmount float64   `json:"transaction_amount"`
}

type SubscriptionToProfile struct {
	ID              string    `json:"id"`
	Status          string    `json:"status"`
	DateCreated     time.Time `json:"date_created"`
	NextPaymentDate time.Time `json:"next_payment_date"`
	Reason          string    `json:"reason"`
}

type SubscriptionToDB struct {
	SubscriptionID    string    `json:"id"`
	UserID            int       `json:"user_id"`
	DateCreated       time.Time `json:"date_created"`
	NextPaymentDate   time.Time `json:"next_payment_date"`
	PreapprovalPlanID string    `json:"plan_id"`
	Status            string    `json:"status"`
}

type UpdateReq struct {
	SubscriptionID  string    `json:"id"`
	NextPaymentDate time.Time `json:"next_payment_date"`
	Status          string    `json:"status"`
}
