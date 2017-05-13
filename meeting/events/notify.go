package events

import (
	"errors"
)

type Notify struct {
	Message string `json:"message"`
}

func (event Notify) Execute() error {
	return errors.New("Not yet implemented!")
}
