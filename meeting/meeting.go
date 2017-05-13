package meeting

import (
	"time"

	"github.com/WeisswurstSystems/WWM-BB/product"
)

type Meeting struct {
	Place    string            `json:"place"`
	Creator  string            `json:"creator"`
	Buyer    string            `json:"buyer"`
	Date     time.Time         `json:"date"`
	Orders   []Order           `json:"orders"`
	Products []product.Product `json:"products"`
}
