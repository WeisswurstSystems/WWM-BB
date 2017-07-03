package user_test

import (
	"encoding/json"
	"testing"

	"github.com/WeisswurstSystems/WWM-BB/user"
)

const userJSON = `{
  "mail": "fabiwilms@gmail.com",
  "roles": [
    "admin",
    "user"
  ],
  "defaultOrders": {
    "Brezen": 1,
    "Weisswurst": 2
  },
  "mailEnabled": true
}`

func TestUserEntity(t *testing.T) {
	testMap := make(map[string]int)
	testMap["Weisswurst"] = 2
	testMap["Brezen"] = 1
	u := user.User{"fabiwilms@gmail.com", "testpassword", []string{"admin", "user"}, testMap, true}
	data, _ := json.MarshalIndent(u, "", "  ")
	if string(data) != userJSON {
		t.Error("Wrong JSON!")
	}
}
