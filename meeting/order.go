package meeting

type CustomerMail string

type OrderItem struct {
	ItemName ProductName `json:"itemName"`
	Amount   int         `json:"amount"`
}

type Order struct {
	Customer CustomerMail      `json:"customer"`
	Payed    bool        `json:"payed"`
	Items    []OrderItem `json:"items"`
}
