package wwm

import (
	"encoding/json"
	"io"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err != nil {
		if ee, ok := err.(*Error); ok {
			http.Error(w, ee.Message, ee.Code)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
}

var (
	ErrNoBody = &Error{"No body was sent", http.StatusBadRequest}
)

func DecodeBody(body io.Reader, v interface{}) error {
	if body == nil {
		return ErrNoBody
	}
	err := json.NewDecoder(body).Decode(&v)
	if err == io.EOF {
		return ErrNoBody
	}
	return err
}
