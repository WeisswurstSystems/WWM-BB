package events

import (
	"errors"
)

type Pay struct {
	UserID string `json:"userID"`
}

func (event Pay) Execute() error {
	return errors.New("Not yet implemented!")
}
