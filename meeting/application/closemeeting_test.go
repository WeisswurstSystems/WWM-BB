package application

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"testing"
)

const ID = "1"

func TestCloseMeeting(t *testing.T) {
	i := NewMockInteractor()

	// GIVEN
	err := i.MeetingStore.Save(meeting.Meeting{ID: ID})
	if err != nil {
		t.Fatal(err)
	}

	err = i.CloseMeeting(CloseMeeting{ID})
	if err != nil {
		t.Fatal(err)
	}

	// THEN
	m, err := i.MeetingStore.FindOne(ID)
	if err != nil {
		t.Fatal(err)
	}
	if !m.Closed {
		t.Fatalf("Meeting was not closed")
	}
}
