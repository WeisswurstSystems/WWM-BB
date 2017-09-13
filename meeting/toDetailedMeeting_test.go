package meeting

import (
	"github.com/WeisswurstSystems/WWM-BB/user"
	"testing"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func TestToDetailedMeeting(t *testing.T) {
	uLogin := user.Login{Mail: "hans@test.de"}
	uPaypal := user.PayPal{"https://paypal.me/hanstest/"}
	testUser := user.User{Login: uLogin, PayPal: uPaypal, Roles: []string{"admin"}}

	testOffer := Offer{Product{"Weisswurst", 0.65},
		Product{"Brezen", 1.30},
		Product{"Weißbier", 2.50}}

	testOrders := []Order{
		{Customer:"A", Items: []OrderItem{{"Brezen", 2}, {"Weißbier", 1}}},
		{Customer:"B", Items: []OrderItem{{"Weisswurst", 1}, {"Brezen", 1}, {"Weißbier", 0}}},
		{Customer:"C", Items: []OrderItem{{"Weisswurst", 6}}},
		{Customer:"D", Items: []OrderItem{{"Weißbier", 2}}},
		{Customer:"E", Items: []OrderItem{{"Brezen", 5}}},
		}

	testMeeting := Meeting{
		Creator: uLogin.Mail,
		Offer:   testOffer,
		Orders:  testOrders,
	}

	resultMeeting := ToDetailedMeeting(testMeeting, testUser.PayPal.MeLink)

	for _, order := range resultMeeting.Orders {

		if order.Customer == "A" {
			assert.Equal(t, order.TotalPrice,5.1, "")
			assert.Equal(t, order.PayLink, uPaypal.MeLink + strconv.FormatFloat(float64(order.TotalPrice), 'f', 2, 32))
		} else if order.Customer == "B" {
			assert.Equal(t, order.TotalPrice, 1.95, "")
		} else if order.Customer == "C" {
			assert.Equal(t, order.TotalPrice, 3.9, "")
		} else if order.Customer == "D" {
			assert.Equal(t, order.TotalPrice, 5.0, "")
		} else if order.Customer == "E" {
			assert.Equal(t, order.TotalPrice, 6.5, "")
		} else {
			t.Fatalf("No Customer is named %v", order.Customer)
		}
	}
	assert.Equal(t, resultMeeting.TotalPrice,22.45, "")
}