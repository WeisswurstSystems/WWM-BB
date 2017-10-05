package user

import (
	"reflect"
	"testing"
)

func TestUser_PublicUser(t *testing.T) {
	tests := []struct {
		name  string
		u     User
		wantP PublicUser
	}{
		{
			name: "shared information of a user",
			u: User{
				Login:  Login{Mail: "peter7995@gmail.com"},
				PayPal: PayPal{MeLink: "asdf"},
				Roles:  []string{"admin"},
			},
			wantP: PublicUser{
				Mail:         "peter7995@gmail.com",
				PayPalMeLink: "asdf",
				Roles:        []string{"admin"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotP := tt.u.PublicUser(); !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("User.PublicUser() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}
