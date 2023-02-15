package models

import "time"

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
