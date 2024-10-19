package types

type Product struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"        validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price"       validate:"required"`
}