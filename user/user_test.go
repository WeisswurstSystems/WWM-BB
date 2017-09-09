package user

import (
	"testing"
)

func TestRegisterted(t *testing.T) {
	var u User
	if !u.IsRegistered() {
		t.Error("User with no registration id should be seen as registered")
	}
	u.RegistrationID = "asdf"
	if u.IsRegistered() {
		t.Error("User should not count as registered")
	}
}
