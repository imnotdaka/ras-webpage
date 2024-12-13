package subscription

import (
	"context"
	"database/sql"
	"time"

	"github.com/imnotdaka/RAS-webpage/internal/database"
)

type Repository struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) CreateSubscriptionToDB(ctx context.Context, req SubscriptionToDB) error {
	_, err := r.DB.ExecContext(ctx, database.CreateSubscriptionQuery, req.SubscriptionID, req.UserID, req.DateCreated, req.NextPaymentDate, req.PreapprovalPlanID, req.Status)
	if err != nil {
		return nil
	}
	return nil
}

func (r Repository) UpdateSubscription(ctx context.Context, req SubscriptionToDB) error {
	_, err := r.DB.ExecContext(ctx, database.UpdateSubscriptionQuery, req.SubscriptionID, req.UserID, req.DateCreated, req.NextPaymentDate, req.PreapprovalPlanID, req.Status, time.Now())
	if err != nil {
		return err
	}
	return nil
}
