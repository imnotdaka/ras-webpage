package plan

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/imnotdaka/RAS-webpage/internal/database"
	"github.com/mercadopago/sdk-go/pkg/preapproval"
)

type repository struct {
	DB *sql.DB
}

type Repository interface {
	CreatePlanDB(ctx context.Context, id string, reason string, frequency int, frequencyType string, transactionAmount float64) (int64, error)
	GetAllPlan(ctx context.Context) ([]PreApprovalPlan, error)
	GetPlanById(ctx context.Context, id string) (preapproval.Request, error)
}

func NewRepo(db *sql.DB) Repository {
	return &repository{DB: db}
}

func (r repository) CreatePlanDB(ctx context.Context, id string, reason string, frequency int, frequencyType string, transactionAmount float64) (int64, error) {
	res, err := r.DB.ExecContext(ctx, database.CreatePlanQuery, id, reason, frequency, frequencyType, transactionAmount)
	if err != nil {
		return 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (r repository) GetAllPlan(ctx context.Context) ([]PreApprovalPlan, error) {
	rows, err := r.DB.QueryContext(ctx, database.GetPlanQuery)
	var plans []PreApprovalPlan
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		plan := PreApprovalPlan{}
		err := rows.Scan(&plan.ID, &plan.Reason, &plan.AutoRecurring.Frequency, &plan.AutoRecurring.FrequencyType, &plan.AutoRecurring.TransactionAmount)
		if err != nil {
			return nil, err
		}
		plans = append(plans, plan)
	}
	fmt.Println(plans)
	return plans, nil
}

func (r repository) GetPlanById(ctx context.Context, id string) (preapproval.Request, error) {
	row := r.DB.QueryRowContext(ctx, database.GetPlanByIdQuery, id)
	plan := preapproval.Request{}
	autorecurring := preapproval.AutoRecurringRequest{}
	err := row.Scan(&plan.PreapprovalPlanID, &plan.Reason, &autorecurring.Frequency, &autorecurring.FrequencyType, &autorecurring.TransactionAmount)
	if err != nil {
		return preapproval.Request{}, err
	}
	plan.AutoRecurring = &autorecurring
	return plan, nil
}
