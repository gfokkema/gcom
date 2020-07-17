package gcommerce

import "time"

type Article struct {
	ID        int64     `db:"id"          json:"id"`
	Title     string    `db:"title"       json:"title"`
	Desc      string    `db:"description" json:"desc"`
	Price     float32   `db:"price"       json:"price"`
	CreatedAt time.Time `db:"created_at"  json:"created_at"`
}
