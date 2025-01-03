package session

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/imnotdaka/RAS-webpage/internal/database"
)

type Session struct {
	UserID  int    `json:"user_id"`
	IsValid bool   `json:"is_valid"`
	Token   string `json:"refresh_token"`
}

type Repository interface {
	Get(ctx context.Context, token string) (Session, error)
	Create(s Session) error
	Update(s Session) error
}

type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, token string) (Session, error) {
	res := r.db.QueryRowContext(ctx, database.GetSessionQuery, token)

	s := Session{}
	err := res.Scan(&s.Token, &s.IsValid)
	if err != nil {
		fmt.Println("err in Get of session scan", err)
		return Session{}, err
	}
	return s, nil
}

func (r *repository) Create(s Session) error {
	_, err := r.db.Exec(database.CreateSessionQuery, s.UserID, s.Token)
	if err != nil {
		fmt.Println("err in Create exec", err)
		return err
	}
	return nil
}

func (r *repository) Update(s Session) error {
	fmt.Println(s.Token, s.IsValid)
	res, err := r.db.Exec(database.UpdateSessionQuery, s.IsValid, s.Token)
	if err != nil {
		fmt.Println("err in Update exec", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		fmt.Println("err in Update rows affected", err)
		return err
	}
	if rows == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}
