package dto

type CategoryDto struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type CategoryDetailDto struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	ProductList []ProductDto `json:"products"`
}
