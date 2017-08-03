package adapter

import (
	"encoding/json"
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/go-errors/errors"
	"github.com/gorilla/mux"
	"net/http"
)

func (ch *QueryHandler) FindAll(w http.ResponseWriter, req *http.Request) error {
	results, err := ch.MeetingStore.FindAll()
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(results)
}

func (ch *QueryHandler) FindByID(w http.ResponseWriter, req *http.Request) error {
	id, ok := mux.Vars(req)["meetingId"]
	if !ok {
		return errors.New("meeting id url parameter missing")
	}

	results, err := ch.MeetingStore.FindOne(meeting.MeetingID(id))
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(results)
}
