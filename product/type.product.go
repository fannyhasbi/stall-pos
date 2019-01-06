package product

type Product struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
}

type ResponseProduct struct {
	Status int       `json:"status"`
	Data   []Product `json:"data"`
}
