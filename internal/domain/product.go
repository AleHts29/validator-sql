package domain

type Producto struct {
	Id          int      `json:"id"`
	Name        string   `json:"nombre" binding:"required"`
	Quantity    *int     `json:"cantidad"`     // Campo que puede ser nulo
	CodeValue   *string  `json:"codigo_valor"` // Campo que puede ser nulo
	IsPublished bool     `json:"publicado"`
	Expiration  *string  `json:"vencimiento"` // Campo que puede ser nulo
	Price       *float64 `json:"precio"`      // Campo que puede ser nulo
}
