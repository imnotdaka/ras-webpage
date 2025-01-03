package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/imnotdaka/RAS-webpage/internal/database"
)

type Repository struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) CreateUser(ctx context.Context, user *User) (int64, error) {
	res, err := r.DB.ExecContext(ctx, database.CreateUserQuery, user.FirstName, user.LastName, user.Email, user.EncryptedPassword)
	if err != nil {
		return 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (r Repository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := r.DB.QueryRowContext(ctx, database.GetUserByEmailQuery, email)
	u := User{}
	err := row.Scan(&u.ID, &u.EncryptedPassword)
	if err != nil {
		fmt.Println(err)
	}
	return u, nil
}

func (r Repository) GetUserById(ctx context.Context, id int) (User, error) {
	fmt.Println(id)
	row := r.DB.QueryRowContext(ctx, database.GetUserByIDQuery, id)
	u := User{}
	err := row.Scan(&u.FirstName, &u.LastName, &u.Email, &u.Subscribed)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r Repository) UpdateUser() (any, error) {
	return nil, nil
}

func (r Repository) DeleteUser(ctx context.Context, id string) (string, error) {
	res, err := r.DB.ExecContext(ctx, database.DeleteUserByIDQuery, id)
	if err != nil {
		return "", err
	}
	affRows, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if affRows == 0 {
		return "", errors.New("no rows deleted")
	}

	return id, nil
}
