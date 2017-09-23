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

// OrderItems is a collection of OrderItems
type OrderItems []OrderItem

// Order contains the order of a single customer.
type Order struct {
	Customer CustomerMail `json:"customer"`
	Payed    bool         `json:"payed"`
	Items    OrderItems   `json:"items"`
}

// Orders is a collection of orders
type Orders []Order

var (
	// ErrOrderNotPresent if the order is not contained in the orders
	ErrOrderNotPresent = wwm.Error{
		Message: "Order of Customer is not contained in the Orders",
		Code:    http.StatusNotFound,
	}
)

// AddItem to the collection of OrderItem. If the product already exists, the amount is added.
func (ois OrderItems) AddItem(item OrderItem) OrderItems {
	i, _, found := ois.FindItemByProductName(item.ItemName)
	if !found {
		return append(ois, item)
	}
	ois[i].Amount += item.Amount
	return ois
}

// FindItemByProductName if it already exists in the collection.
func (ois OrderItems) FindItemByProductName(name ProductName) (index int, item OrderItem, found bool) {
	for i, item := range ois {
		if item.ItemName == name {
			return i, item, true
		}
	}
	return -1, OrderItem{}, false
}

// FindOrderByCustomer in the order collection. If not found return a new Order for the customer.
func (orders Orders) FindOrderByCustomer(Customer CustomerMail) (index int, order Order, found bool) {
	for i, order := range orders {
		if order.Customer == Customer {
			return i, order, true
		}
	}
	return -1, Order{}, false
}

// AddOrderItemForCustomer in the order collection. If no order for the customer exists, a new one is created.
func (orders Orders) AddOrderItemForCustomer(item OrderItem, Customer CustomerMail) Orders {
	i, order, found := orders.FindOrderByCustomer(Customer)
	if !found {
		order.Items = []OrderItem{item}
		order.Customer = Customer
		return append(orders, order)
	}
	orders[i].Items = order.Items.AddItem(item)
	return orders
}
