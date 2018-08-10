package product_test

import (
	"encoding/json"
	"testing"

	"github.com/WeisswurstSystems/WWM-BB/product"
)

const productJSON = `{
  "name": "Weisswurst",
  "price": 120
}`

func TestProductEntity(t *testing.T) {
	p := product.Product{"Weisswurst", 120}
	data, _ := json.MarshalIndent(p, "", "  ")
	if string(data) != productJSON {
		t.Error("Wrong JSON!")
	}
}
