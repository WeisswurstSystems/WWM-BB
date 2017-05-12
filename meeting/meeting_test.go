package meeting_test

import (
	"encoding/json"
	"testing"

	"github.com/WeisswurstSystems/WWM-BB/meeting"
)

const meetingJSON = `{
  "date": "0001-01-01T00:00:00Z",
  "orders": [
    {
      "customer": "peter-mueller@github.com",
      "items": [
        {
          "name": "Weisswurst",
          "amount": 3
        },
        {
          "name": "Brezen",
          "amount": 2
        }
      ]
    }
  ]
}`

func TestMeetingEntity(t *testing.T) {
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
