package model

import "time"

type OderDetail struct {
	Order_id      int64
	Product_id    int64
	Product_name  string
	Product_price float64
	Created_at    time.Time
}
