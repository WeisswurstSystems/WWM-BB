package event

import "fmt"

type EventError struct {
	Message string
	Code    int
}

func (err EventError) Error() string {
	return fmt.Sprintf("%v: %v", err.Code, err.Message)
}
