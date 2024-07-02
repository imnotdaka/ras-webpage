package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/imnotdaka/RAS-webpage/cmd/config"
)

func NewDB(dbcfg config.DB) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbcfg.User, dbcfg.Password, dbcfg.Ip, dbcfg.Port, dbcfg.Database))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
