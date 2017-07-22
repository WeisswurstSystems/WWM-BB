package adapter

import (
	"github.com/WeisswurstSystems/WWM-BB/database"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

const singleUserJson = `[{"mail":"fabian.wilms@gmail.com","roles":[],"defaultOrders":{},"mailEnabled":false}]
`

var qh QueryHandler

func TestRead(t *testing.T) {
	database.Users.DropCollection()
	store.Save(user.User{Mail: "fabian.wilms@gmail.com"})

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h := http.HandlerFunc(qh.Read)
	h.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() != singleUserJson {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), singleUserJson)
	}
}
