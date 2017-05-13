package events

import (
	"errors"
)

type AssignBuyer struct {
	Buyer string `json:"buyer"`
}

func (event AssignBuyer) Execute() error {
	return errors.New("Not yet implemented!")
}
