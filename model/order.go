package model

import "time"

type Order struct {
	Id           int64     `db:"id" json:"id"`
	ClientId     int64     `db:"client_id" json:"client_id"`
	OrderedAt    string    `db:"ordered_at" json:"ordered_at"`
	TotalPrice   float64   `db:"total_price" json:"total_price"`
	DeliveryDate string    `db:"delivery_date" json:"delivery_date"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Description  string    `db:"description" json:"description"`
}
