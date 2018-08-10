package meeting_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/product"
)

const meetingJSON = `{
  "id": "123459876",
  "place": "somewhere over the rainbow",
  "creator": "fabiwilms@gmail.com",
  "buyer": "fabiwilms@gmail.com",
  "date": "0001-01-01T00:00:00Z",
  "closeDate": "0001-01-01T00:00:00Z",
  "closed": false,
  "orders": [
    {
      "customer": "peter-mueller@github.com",
      "payed": false,
      "items": [
        {
          "itemName": "Weisswurst",
          "amount": 3
        },
        {
          "itemName": "Brezen",
          "amount": 2
        }
      ]
    }
  ],
  "products": [
    {
      "name": "Weisswurst",
      "price": 1.05
    },
    {
      "name": "Brezen",
      "price": 0.63
    }
  ]
}`

func TestMeetingEntity(t *testing.T) {
	items := []meeting.OrderItem{{"Weisswurst", 3}, {"Brezen", 2}}
	dateTime, _ := time.Parse("2014-09-12T11:45:26.371Z", "0001-01-01T00:00:00Z")
	m := meeting.Meeting{
		ID:        "123459876",
		Place:     "somewhere over the rainbow",
		Creator:   "fabiwilms@gmail.com",
		Buyer:     "fabiwilms@gmail.com",
		Date:      dateTime,
		CloseDate: dateTime,
		Closed:    false,
		Orders: []meeting.Order{
			{"peter-mueller@github.com", false, items},
		},
		Products: []product.Product{{"Weisswurst", 1.05}, {"Brezen", .63}},
	}

	data, _ := json.MarshalIndent(m, "", "  ")
	if string(data) != meetingJSON {
		t.Error("JSON does not match!")
	}
}
