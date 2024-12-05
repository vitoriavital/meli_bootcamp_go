package model

type Product struct {
	Id			int		`json:"id"`
	Name		string	`json:"name"`
	Quantity	int		`json:"quantity"`
	CodeValue	string	`json:"code_value"`
	IsPublished	bool	`json:"is_published"`
	Expiration	string	`json:"expiration"`
	Price		float64	`json:"price"`
}

type RequestBodyProduct struct {
	Name		string	`json:"name"`
	Quantity	int		`json:"quantity"`
	CodeValue	string	`json:"code_value"`
	IsPublished	bool	`json:"is_published"`
	Expiration	string	`json:"expiration"`
	Price		float64	`json:"price"`
}

type ResponseBodyProduct struct {
	Message		string			`json:"message"`
	Product	*Product 	`json:"product,omitempty"`
	Error		bool			`json:"error"`
}
