package models

import (
	"armageddon/internal/validator"
	"context"
	"database/sql"
	"errors"
	"time"
)

type CarModel struct {
	DB *sql.DB
}

type Car struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Brand       string    `json:"brand"`
	Description string    `json:"description"`
	Color       string    `json:"color,omitempty"`
	Year        int32     `json:"year,omitempty"`
	Price       int32     `json:"price"`
	IsUsed      bool      `json:"is_used"`
}

func (m CarModel) Insert(car *Car) error {
	query := `
		INSERT INTO car (brand, description, color, year, price)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, is_used`

	args := []any{car.Brand, car.Description, car.Color, car.Year, car.Price}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&car.ID, &car.CreatedAt, &car.IsUsed)
}

func (m CarModel) Get(id int64) (*Car, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `SELECT * FROM car WHERE id = $1`

	var car Car

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&car.ID,
		&car.CreatedAt,
		&car.Brand,
		&car.Description,
		&car.Color,
		&car.Year,
		&car.Price,
		&car.IsUsed,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &car, nil
}

func (m CarModel) Update(car *Car) error {
	query := `
		UPDATE car
		SET brand = $1, description = $2, color = $3, year = $4, price = $5, is_used = $6
		WHERE id = $7
		RETURNING created_at`

	args := []any{
		car.Brand,
		car.Description,
		car.Color,
		car.Year,
		car.Price,
		car.IsUsed,
		car.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&car.CreatedAt)
}

func (m CarModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `DELETE FROM car WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func ValidateCar(v *validator.Validator, car *Car) {
	v.Check(car.Brand != "", "brand", "must be provided")
	v.Check(len(car.Brand) <= 500, "brand", "must not be more than 500 bytes long")

	v.Check(car.Description != "", "brand", "must be provided")
	v.Check(car.Color != "", "color", "must be provided")

	v.Check(car.Year != 0, "year", "must be provided")
	v.Check(car.Year >= 1888, "year", "must be greater than 1888")
	v.Check(car.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(car.Price != 0, "price", "must be provided")
	v.Check(car.Price > 0, "price", "must be a positive integer")
}
