package model

import "time"

type Product struct {
	Id          uint      `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	IsAvailable bool      `db:"is_available"`
	Price       float64   `db:"price"`
	QtdStored   int64     `db:"qtd_stored"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
