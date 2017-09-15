package closemeeting

import (
	"github.com/WeisswurstSystems/WWM-BB/user"
	"testing"
)

func TestCloseMeeting(t *testing.T) {
	u, mock := NewMocking()
	mock.Meeting.ID = "1"

	err := u.CloseMeeting(Request{"1", user.Login{}})
	if err != nil {
		t.Fatal(err)
	}
	if !mock.Saved.Closed {
		t.Fatalf("Meeting was not closed")
	}
}
