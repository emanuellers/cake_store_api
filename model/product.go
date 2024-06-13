package model

import "time"

type Product struct {
	Id          uint      `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	IsAvailable bool      `db:"is_available" json:"is_available"`
	Price       float64   `db:"price" json:"price"`
	QtdStored   int64     `db:"qtd_stored" json:"qtd_stored"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
