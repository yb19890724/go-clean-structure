package json

import "time"

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Price       float32   `json:"price"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}
