package query

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/driver"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
	"github.com/WeisswurstSystems/WWM-BB/wwm"
)

func TestQueryHandler_FindAll(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	login := user.Hans
	req.SetBasicAuth(login.Mail, login.Password)

	store := driver.NewMemoryStore()
	store.Save(user.Hans)

	q := QueryHandler{
		Store:               store,
		AuthenticateUseCase: authenticate.Interactor{ReadStore: store},
	}

	rr := httptest.NewRecorder()
	handler := wwm.Handler(q.FindAll)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"mail":"hans@gmail.com","payPalMeLink":"http://www.paypal.me/hans","roles":["admin"]}]
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestQueryHandler_Identity(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	login := user.Hans
	req.SetBasicAuth(login.Mail, login.Password)

	store := driver.NewMemoryStore()
	store.Save(user.Hans)

	q := QueryHandler{
		Store:               store,
		AuthenticateUseCase: authenticate.Interactor{ReadStore: store},
	}

	rr := httptest.NewRecorder()
	handler := wwm.Handler(q.Identity)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"login":{"mail":"hans@gmail.com","password":"hansistsogeil"},"payPal":{"meLink":"http://www.paypal.me/hans"},"registrationID":"","roles":["admin"],"defaultOrders":null,"mailEnabled":false}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
