package models

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Car   CarModel
	Users UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Car:   CarModel{DB: db},
		Users: UserModel{DB: db},
	}
}
