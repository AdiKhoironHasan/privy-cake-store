package models

type CakeModels struct {
	ID          int     `db:"id"`
	Title       string  `db:"title"`
	Description string  `db:"description"`
	Rating      float64 `db:"rating"`
	Image       string  `db:"image"`
	Created_at  string  `db:"created_at"`
	Updated_at  string  `db:"updated_at"`
}
