package mail

import (
	"bytes"
	"log"
	"text/template"
	"github.com/WeisswurstSystems/WWM-BB/util"
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

func NewContent(data SmtpTemplateData) ([]byte, error) {
	var doc bytes.Buffer
	context := SmtpTemplateData{
		getButtonDisplay(data),
		data.BodyButtonLink,
		data.BodyButtonText,
		data.Subject,
		data.BodyShortText,
		data.Body,
		util.GetUID(8),
	}
	err := t.Execute(&doc, &context)
	if err != nil {
		log.Printf("%v error trying to execute mail template", LOG_TAG)
		return nil, err
	}
	return doc.Bytes(), nil
}

/**
 * Wenn kein Link + Text angegeben wird, soll der Button in der Mail nicht dargestellt werden.
 */
func getButtonDisplay(data SmtpTemplateData) string {
	if data.BodyButtonLink != "" && data.BodyButtonText != "" {
		return "block"
	} else {
		return "none"
	}
}
