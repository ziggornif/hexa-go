package todo

import "time"

// Todo - todo model
type Todo struct {
	ID        string    `pg:",pk" json:"id"`
	Title     string    `pg:",unique" json:"title"`
	Completed bool      `json:"completed"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
