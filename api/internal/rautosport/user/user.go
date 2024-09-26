package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/imnotdaka/RAS-webpage/internal/database"
)

type Repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return Repository{db: db}
}

func (r Repository) CreateUser(user *User) (int64, error) {
	res, err := r.db.Exec(database.CreateUserQuery, user.FirstName, user.LastName, user.Email, user.EncryptedPassword)
	if err != nil {
		return 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (r Repository) GetUserByEmail(email string) (User, error) {
	row := r.db.QueryRow(database.GetUserByEmailQuery, email)
	u := User{}
	err := row.Scan(&u.ID, &u.EncryptedPassword)
	if err != nil {
		fmt.Println(err)
	}
	return u, nil
}

func (r Repository) GetUserById(id int) (User, error) {
	fmt.Println(id)
	row := r.db.QueryRow(database.GetUserByIDQuery, id)
	u := User{}
	err := row.Scan(&u.FirstName, &u.LastName, &u.Email)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r Repository) UpdateUser() (any, error) {
	return nil, nil
}

func (r Repository) DeleteUser(id string) (string, error) {
	res, err := r.db.Exec(database.DeleteUserByIDQuery, id)
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
