package webhandler

import (
	"bytes"
	"encoding/json"
	"github.com/WeisswurstSystems/WWM-BB/database"
	"github.com/WeisswurstSystems/WWM-BB/mail"
	"github.com/WeisswurstSystems/WWM-BB/user/event"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var ch CommandHandler

func TestRegister(t *testing.T) {
	database.Users.DropCollection()

	event := event.Register{
		Mail:        "p.mueller",
		Password:    "password",
		MailEnabled: true,
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&event)
	req, err := http.NewRequest("POST", "/", &buf)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h := http.HandlerFunc(ch.Register)
	h.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	m := mail.Mail{} //TODO

	if len(m.Receivers) != 1 {
		t.Errorf("Contained more than one Receiver: %v", m.Receivers)
	}
	if m.Receivers[0] != event.Mail {
		t.Errorf("Receiver is wrong: got %v want %v", m.Receivers[0], event.Mail)
	}

	if !strings.Contains(string(m.Content), "Subject: Deine Registrierung bei der Weisswurst-Verwaltung") {
		t.Error("Content did not contain the right subject")
	}
}
