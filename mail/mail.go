package mail

import (
	"bytes"
	"log"
	"text/template"
)

const LOG_TAG = "[MAIL_CLIENT]"

type Service interface {
	Send(mail Mail) error
}

type Mail struct {
	Content   []byte
	Receivers []string
}

var t = template.New("EmailTemplate")

func init() {
	_, err := t.Parse(EmailTemplate)
	if err != nil {
		log.Fatalf("%v error trying to parse mail template", LOG_TAG)
	}
}

func NewContent(subject string, message string) ([]byte, error) {
	var doc bytes.Buffer
	context := SmtpTemplateData{subject, message}
	err := t.Execute(&doc, &context)
	if err != nil {
		log.Printf("%v error trying to execute mail template", LOG_TAG)
		return nil, err
	}
	return doc.Bytes(), nil
}
