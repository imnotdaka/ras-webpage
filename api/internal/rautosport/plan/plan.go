package plan

import (
	"database/sql"
	"fmt"

	"github.com/imnotdaka/RAS-webpage/internal/database"
)

type repository struct {
	DB *sql.DB
}

type Repository interface {
	CreatePlanDB(*string, *string, *int, *string, *int) (int64, error)
	GetAll() ([]PreApprovalPlan, error)
}

func NewRepo(db *sql.DB) Repository {
	return &repository{DB: db}
}

func (r repository) CreatePlanDB(id *string, reason *string, frequency *int, frequencyType *string, transactionAmount *int) (int64, error) {
	res, err := r.DB.Exec(database.CreatePlanQuery, id, reason, frequency, frequencyType, transactionAmount)
	if err != nil {
		return 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (r repository) GetAll() ([]PreApprovalPlan, error) {
	rows, err := r.DB.Query(database.GetPlanQuery)
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
