package meeting

type Offer []Product
type ProductName string

type Product struct {
	Name  ProductName `json:"name"`
	Price float64     `json:"price"`
}

func (o Offer) Put(p Product) Offer {
	i, found := o.indexOf(p.Name)
	if found {
		o[i] = p
	} else {
		o = append(o, p)
	}
	return o
}

func (o Offer) Remove(name ProductName) Offer {
	i, found := o.indexOf(name)
	if found {
		o = append(o[:i], o[i+1:]...)
	}
	return o
}

func (o Offer) indexOf(name ProductName) (pos int, found bool) {
	for i, p := range o {
		if p.Name == name {
			return i, true
		}
	}
	return -1, false
}
