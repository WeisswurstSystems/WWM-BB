package meeting_test

import (
	"encoding/json"

	"github.com/WeisswurstSystems/WWM-BB/meeting"

	"testing"
)

const meetingJSON = `{
  "Date": "0001-01-01T00:00:00Z",
  "Orders": [
    {
      "Customer": "peter-mueller@github.com",
      "Items": [
        {
          "Name": "Weisswurst",
          "Amount": 3
        },
        {
          "Name": "Brezen",
          "Amount": 2
        }
      ]
    }
  ]
}`

func TestEntityOK(t *testing.T) {

	items := []meeting.OrderItem{{"Weisswurst", 3}, {"Brezen", 2}}
	m := meeting.Meeting{
		Orders: []meeting.Order{
			{"peter-mueller@github.com", items},
		},
	}
	data, _ := json.MarshalIndent(m, "", "  ")

	if string(data) != meetingJSON {
		t.Error("JSON does not match!")
	}
}
