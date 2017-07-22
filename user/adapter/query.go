package adapter

import (
	"encoding/json"
	"net/http"
)

func (ch *QueryHandler) Read(w http.ResponseWriter, req *http.Request) {
	results, err := ch.UserStore.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
