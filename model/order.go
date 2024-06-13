package model

import "time"

type Order struct {
	Id           int64     `db:"id"`
	ClientId     int64     `db:"client_id"`
	OrderedAt    time.Time `db:"ordered_at"`
	Total_Price  float64   `db:"total_price"`
	DeliveryDate time.Time `db:"delivery_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	Description  string    `db:"description"`
}
