package handler

import (
	"encoding/json"
	"github.com/WeisswurstSystems/WWM-BB/user/store"
	"net/http"
)

func Read(w http.ResponseWriter, req *http.Request) {
	results, err := store.FindAll()
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
