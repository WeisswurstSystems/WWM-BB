package meeting

import (
	"time"

	"github.com/WeisswurstSystems/WWM-BB/product"
)

type Meeting struct {
	ID        string            `json:"id"`
	Place     string            `json:"place"`
	Creator   string            `json:"creator"`
	Buyer     string            `json:"buyer"`
	Date      time.Time         `json:"date"`
	CloseDate time.Time         `json:"closeDate"`
	Closed    bool              `json:"closed"`
	Orders    []Order           `json:"orders"`
	Products  []product.Product `json:"products"`
}

type ReducedMeeting struct {
	ID        string            `json:"id"`
	Place     string            `json:"place"`
	Date      time.Time         `json:"date"`
	CloseDate time.Time         `json:"closeDate"`
	Closed    bool              `json:"closed"`
}