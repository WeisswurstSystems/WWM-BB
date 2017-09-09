package query

import (
	"encoding/json"
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/go-errors/errors"
	"github.com/gorilla/mux"
	"net/http"
)

type QueryHandler struct {
	MeetingStore meeting.ReadStore
}

func (ch *QueryHandler) FindAll(w http.ResponseWriter, req *http.Request) error {
	results, err := ch.MeetingStore.FindAll()
	if err != nil {
		return err
	}

	reduced := meeting.AllReduced(results)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(reduced)
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
