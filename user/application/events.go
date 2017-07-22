package application

import "net/http"

type Activate struct {
	RegistrationID string `json:"registrationID"`
}

type Register struct {
	Mail        string `json:"mail"`
	Password    string `json:"password"`
	MailEnabled bool   `json:"mailEnabled"`
}

func (e Register) Validate() error {
	if e.Mail == "" {
		return &Error{"Missing filed: mail", http.StatusBadRequest}
	}
	if e.Password == "" {
		return &Error{"Missing filed: password", http.StatusBadRequest}
	}
	return nil
}
