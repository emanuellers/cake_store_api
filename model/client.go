package model

type Client struct {
	FirstName string `db:"firstname" json:"firstname"`
	LastName  string `db:"lastname" json:"lastname"`
	Email     string `db:"email" json:"email"`
	Id        uint   `db:"id" json:"id"`
}
