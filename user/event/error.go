package event

import "fmt"

type Error struct {
	Message string
	Code    int
}

func (err Error) Error() string {
	return fmt.Sprintf("%v: %v", err.Code, err.Message)
}
