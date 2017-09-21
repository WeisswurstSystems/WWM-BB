package meeting

// Offer is a list of products that can be bought.
type Offer []Product

// ProductName identifies a product by its name.
type ProductName string

// Product that has a price.
type Product struct {
	Name  ProductName `json:"name"`
	Price float64     `json:"price"`
}

// Put the product into the offer, if it does not already exist.
func (o Offer) Put(p Product) Offer {
	i, found := o.indexOf(p.Name)
	if found {
		o[i] = p
	} else {
		o = append(o, p)
	}
	return o
}

// Remove a Product by name from an offer.
func (o Offer) Remove(name ProductName) Offer {
	i, found := o.indexOf(name)
	if found {
		o = append(o[:i], o[i+1:]...)
	}
	return o
}

// indexOf a product (by name) in the offer.
func (o Offer) indexOf(name ProductName) (pos int, found bool) {
	for i, p := range o {
		if p.Name == name {
			return i, true
		}
	}
	return -1, false
}
