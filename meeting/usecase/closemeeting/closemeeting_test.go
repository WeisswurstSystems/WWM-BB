package closemeeting

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"testing"
)

func TestCloseMeeting(t *testing.T) {
	u, mock := NewMocking()
	mock.Meeting.ID = "1"

	err := u.CloseMeeting(Request{"1"})
	if err != nil {
		t.Fatal(err)
	}
	if !mock.Saved.Closed {
		t.Fatalf("Meeting was not closed")
	}
}

func TestIsAllowed(t *testing.T) {
	tests := []struct {
		User    user.User
		Meeting meeting.Meeting
		Allowed bool
	}{
		{
			user.User{Roles: []string{"admin"}},
			meeting.Meeting{},
			true,
		}, {
			user.User{Mail: "test@mail.com"},
			meeting.Meeting{Creator: "test@mail.com"},
			true,
		}, {
			user.User{Mail: "test@mail.com"},
			meeting.Meeting{},
			false,
		},
	}

	for _, test := range tests {
		r := isAllowed(test.User, test.Meeting)
		if r != test.Allowed {
			t.Errorf("isAllowed(%v, %v) => %v, want %v", test.User, test.Meeting, r, test.Allowed)
		}
	}
}
