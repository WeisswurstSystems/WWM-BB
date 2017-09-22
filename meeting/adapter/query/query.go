package query

import (
	"encoding/json"
	"net/http"

	"errors"

	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/gorilla/mux"
)

type QueryHandler struct {
	MeetingStore meeting.ReadStore
	UserStore    user.ReadStore
}

func (ch *QueryHandler) FindAll(w http.ResponseWriter, req *http.Request) error {
	results, err := ch.MeetingStore.FindAll()
	if err != nil {
		return err
	}

	reduced := meeting.AllReduced(results)

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(reduced)
}

func (ch *QueryHandler) FindByID(w http.ResponseWriter, req *http.Request) error {
	id, ok := mux.Vars(req)["meetingId"]
	if !ok {
		return errors.New("meeting id url parameter missing")
	}

	result, err := ch.MeetingStore.FindOne(meeting.MeetingID(id))
	if err != nil {
		return err
	}

	creatorUser, err := ch.UserStore.FindByMail(result.Creator)

	if err != nil {
		return errors.New("Creator wurde nicht in Nutzer-DB gefunden")
	}

	detailedResult := result.Detailed(creatorUser.PayPal.MeLink)

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(detailedResult)
}
