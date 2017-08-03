package events

import (
	"errors"
)

type Create struct {
	Creator string `json:"creator"`
}

func (event Create) Execute() error {
	return errors.New("Not yet implemented!")
}
