package model

import "time"

type Order struct {
	Id            int64
	Client_Id     int64
	Ordered_at    time.Time
	Total_Price   float64
	Delivery_Date time.Time
	Updated_at    time.Time
	Description   string
}
