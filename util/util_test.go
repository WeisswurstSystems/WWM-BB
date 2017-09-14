package util

import (
	"testing"
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/user"
)

func TestIsMeetingCreator(t *testing.T) {
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
			user.User{Login: user.Login{Mail: "test@mail.com"}},
			meeting.Meeting{Creator: "test@mail.com"},
			true,
		}, {
			user.User{Login: user.Login{Mail: "test@mail.com"}},
			meeting.Meeting{},
			false,
		},
	}

	for _, test := range tests {
		r := IsMeetingCreator(test.User, test.Meeting)
		if r != test.Allowed {
			t.Errorf("isAllowed(%v, %v) => %v, want %v", test.User, test.Meeting, r, test.Allowed)
		}
	}
}
