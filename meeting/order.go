package meeting

type OrderItem struct {
	Name   string `json:name`
	Amount int    `json:amount`
}

type Order struct {
	Customer string      `json:customer`
	Items    []OrderItem `json:items`
}
