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

func (r Repository) GetSubscriptionByUserID(ctx context.Context, userID int) (SubscriptionToProfile, error) {
	row := r.DB.QueryRowContext(ctx, database.GetSubscriptionByUserIDQuery, userID)
	s := SubscriptionToProfile{}
	err := row.Scan(&s.ID, &s.Reason, &s.Status, &s.DateCreated, &s.NextPaymentDate)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (r Repository) UpdateSubscription(ctx context.Context, req UpdateReq) error {
	_, err := r.DB.ExecContext(ctx, database.UpdateSubscriptionQuery, req.NextPaymentDate, req.Status, time.Now(), req.SubscriptionID)
	if err != nil {
		return err
	}
	return nil
}
