package meeting

import (
	"net/http"

	"github.com/WeisswurstSystems/WWM-BB/wwm"
)

// CustomerMail for representing the customer by mail.
type CustomerMail string

// OrderItem is a single item of a order.
type OrderItem struct {
	ItemName ProductName `json:"itemName"`
	Amount   int         `json:"amount"`
}

// Order contains the order of a single customer.
type Order struct {
	Customer CustomerMail `json:"customer"`
	Payed    bool         `json:"payed"`
	Items    []OrderItem  `json:"items"`
}

var (
	// ErrOrderNotPresent if the order is not contained in the orders
	ErrOrderNotPresent = wwm.Error{
		Message: "Order of Customer is not contained in the Orders",
		Code:    http.StatusNotFound,
	}
)

// AddItem to the collection of OrderItem. If the product already exists, the amount is added.
func (order *Order) AddItem(item OrderItem) {
	i, _, found := order.FindItemByProductName(item.ItemName)
	if !found {
		order.Items = append(order.Items, item)
		return
	}
	order.Items[i].Amount += item.Amount
}

// FindItemByProductName if it already exists in the collection.
func (order *Order) FindItemByProductName(name ProductName) (index int, item OrderItem, found bool) {
	for i, item := range order.Items {
		if item.ItemName == name {
			return i, item, true
		}
	}
	return -1, OrderItem{}, false
}
