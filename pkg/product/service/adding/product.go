package adding

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Price       float32   `json:"price"`
	Description string    `json:"description"`
}
