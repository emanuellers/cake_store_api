package model

import "time"

type OderDetail struct {
	OrderId      int64     `db:"order_id" json:"order_id"`
	ProductId    int64     `db:"product_id" json:"product_id"`
	ProductName  string    `db:"product_name" json:"product_name"`
	ProductPrice float64   `db:"product_price" json:"product_price"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
