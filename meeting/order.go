package meeting

import (
	"github.com/WeisswurstSystems/WWM-BB/util"
	"strconv"
	"strings"
)

type OrderItem struct {
	ItemName ProductName `json:"itemName"`
	Amount   int         `json:"amount"`
}

type Order struct {
	Customer string      `json:"customer"`
	Payed    bool        `json:"payed"`
	Items    []OrderItem `json:"items"`
}

type DetailedOrder struct {
	Customer   string      `json:"customer"`
	Payed      bool        `json:"payed"`
	Items      []OrderItem `json:"items"`
	TotalPrice float64     `json:"totalPrice"`
	PayLink    string      `json:"payLink"`
}

func ToDetailedOrders(orders []Order, products []Product, payPalLink string) []DetailedOrder {
	var detailedOrders []DetailedOrder
	for _, order := range orders {
		detailedOrders = append(detailedOrders, ToDetailedOrder(order, products, payPalLink))
	}
	return detailedOrders
}

func ToDetailedOrder(order Order, products []Product, paypalLink string) DetailedOrder {
	var detailedOrder = DetailedOrder{
		Customer: order.Customer,
		Payed:    order.Payed,
		Items:    order.Items,
	}

	var resultPrice float64

	for _, item := range detailedOrder.Items {
		price := products[
			util.IndexOf(len(products), func(i int) bool {
				return products[i].Name == item.ItemName
			})].Price
		resultPrice += (float64(item.Amount) * price)
	}

	detailedOrder.TotalPrice = util.FloatToFixed(resultPrice, 2)

	if paypalLink != "" {
		if !strings.HasSuffix(paypalLink, "/") {
			paypalLink += "/"
		}
		detailedOrder.PayLink = paypalLink +  strconv.FormatFloat(float64(detailedOrder.TotalPrice), 'f', 2, 32)
	}

	return detailedOrder
}

