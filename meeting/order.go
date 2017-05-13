package meeting

type OrderItem struct {
	ItemName string `json:"itemName"`
	Amount   int    `json:"amount"`
}

type Order struct {
	Customer string      `json:"customer"`
	Payed    bool        `json:"payed"`
	Items    []OrderItem `json:"items"`
}