package meeting

import (
	"strconv"
	"strings"
	"time"

	"github.com/WeisswurstSystems/WWM-BB/util"
)

type DetailedMeeting struct {
	ID         MeetingID       `json:"id"`
	Place      string          `json:"place"`
	Creator    string          `json:"creator"`
	Buyer      string          `json:"buyer"`
	Date       time.Time       `json:"date"`
	CloseDate  time.Time       `json:"closeDate"`
	Closed     bool            `json:"closed"`
	Orders     []DetailedOrder `json:"orders"`
	Offer      Offer           `json:"offer"`
	TotalPrice float64         `json:"totalPrice"`
	TotalItems []OrderItem     `json:"totalItems"`
}

type DetailedOrder struct {
	Customer   CustomerMail      `json:"customer"`
	Payed      bool        `json:"payed"`
	Items      []OrderItem `json:"items"`
	TotalPrice float64     `json:"totalPrice"`
	PayLink    string      `json:"payLink"`
}

func (m Meeting) Detailed(paypalLink string) DetailedMeeting {
	detailedMeeting := DetailedMeeting{
		ID:        m.ID,
		Place:     m.Place,
		Creator:   m.Creator,
		Buyer:     m.Buyer,
		Date:      m.Date,
		CloseDate: m.CloseDate,
		Closed:    m.Closed,
		Orders:    ToDetailedOrders(m.Orders, m.Offer, paypalLink),
		Offer:     m.Offer,
	}

	var totalPrice float64
	var totalItems []OrderItem

	for _, order := range detailedMeeting.Orders {
		totalPrice += order.TotalPrice
		for _, item := range order.Items {
			index := util.IndexOf(len(totalItems), func(i int) bool {
				return totalItems[i].ItemName == item.ItemName
			})
			if index == -1 {
				totalItems = append(totalItems, OrderItem{
					ItemName: item.ItemName,
					Amount:   item.Amount,
				})
			} else {
				totalItems[index].Amount = totalItems[index].Amount + item.Amount
			}
		}
	}

	detailedMeeting.TotalPrice = totalPrice
	detailedMeeting.TotalItems = totalItems

	return detailedMeeting
}

func ToDetailedOrders(orders []Order, products []Product, payPalLink string) []DetailedOrder {
	var detailedOrders []DetailedOrder
	for _, order := range orders {
		detailedOrders = append(detailedOrders, order.Detailed(products, payPalLink))
	}
	return detailedOrders
}

func (order Order) Detailed(products []Product, paypalLink string) DetailedOrder {
	var detailedOrder = DetailedOrder{
		Customer: order.Customer,
		Payed:    order.Payed,
		Items:    order.Items,
	}

	var resultPrice float64

	for _, item := range detailedOrder.Items {
		price := products[util.IndexOf(len(products), func(i int) bool {
			return products[i].Name == item.ItemName
		})].Price
		resultPrice += (float64(item.Amount) * price)
	}

	detailedOrder.TotalPrice = util.FloatToFixed(resultPrice, 2)

	if paypalLink != "" {
		if !strings.HasSuffix(paypalLink, "/") {
			paypalLink += "/"
		}
		detailedOrder.PayLink = paypalLink + strconv.FormatFloat(float64(detailedOrder.TotalPrice), 'f', 2, 32)
	}

	return detailedOrder
}
