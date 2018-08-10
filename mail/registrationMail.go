package mail

import (
	"text/template"
	"bytes"
)

type registrationData struct{
	Usermail string
	RegistrationID string
}

const topic = "Deine Registrierung bei der Weisswurst-Verwaltung"

const messageTemplate = `
Hallo {{.Usermail}}!

Vielen Dank für deine Registrierung.
Um Dein Konto zu aktivieren, klicke bitte auf den folgenden Link:
http://wwm-bb.herokuapp.com/users/register/{{.RegistrationID}}

Viel Spaß beim bestellen!
`

func SendRegistrationMail(registrationId string, usermail string) {
	data := registrationData{Usermail: usermail, RegistrationID: registrationId}
	tmpl, err := template.New("test").Parse(messageTemplate)
	if err != nil { panic(err) }

	var message bytes.Buffer
	err = tmpl.Execute(&message, data)
	if err != nil { panic(err) }

	SendMail(topic, message.String(), []string{usermail})
}
