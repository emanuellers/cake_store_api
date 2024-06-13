package model

import "time"

type OderDetail struct {
	OrderId      int64     `db:"order_id"`
	ProductId    int64     `db:"product_id"`
	ProductName  string    `db:"product_name"`
	ProductPrice float64   `db:"product_price"`
	CreatedAt    time.Time `db:"created_at"`
}
