package model

import "time"

type Product struct {
	Id           int64
	Name         string
	Description  string
	Is_available bool
	Price        float64
	Qtd_stored   int64
	Created_at   time.Time
	Updated_at   time.Time
}
