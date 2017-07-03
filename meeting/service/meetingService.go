package service

import (
	"encoding/json"
	"net/http"

	"github.com/WeisswurstSystems/WWM-BB/meeting/store"
)

func Read(w http.ResponseWriter, req *http.Request) {
	results, err := store.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var js []byte
	js, err = json.Marshal(results)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
