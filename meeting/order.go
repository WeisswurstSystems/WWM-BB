package meeting

// OrderItem is a single item of a order.
type OrderItem struct {
	ItemName ProductName `json:"itemName"`
	Amount   int         `json:"amount"`
}

// Order contains the order of a single customer.
type Order struct {
	Customer string      `json:"customer"`
	Payed    bool        `json:"payed"`
	Items    []OrderItem `json:"items"`
}
