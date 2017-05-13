package user_test

import (
	"encoding/json"
	"testing"

	"github.com/WeisswurstSystems/WWM-BB/user"
)

const userJSON = `{
  "userID": "12345",
  "mail": "fabiwilms@gmail.com"
}`

func TestUserEntity(t *testing.T) {
	u := user.User{"12345", "fabiwilms@gmail.com"}
	data, _ := json.MarshalIndent(u, "", "  ")
	if string(data) != userJSON {
		t.Error("Wrong JSON!")
	}
}
