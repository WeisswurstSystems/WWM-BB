package adapter

import (
	"encoding/json"
	"net/http"
)

func (ch *QueryHandler) FindAll(w http.ResponseWriter, req *http.Request) error {
	results, err := ch.UserStore.FindAll()
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(results)
}
