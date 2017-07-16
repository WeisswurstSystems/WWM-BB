package mail

import (
	"log"
	"os"
	"net/smtp"
	"text/template"
	"strconv"
	"strings"
	"bytes"
)

type EmailUser struct {
	Username    string
	Password    string
	EmailServer string
	Port        int
}

type SmtpTemplateData struct {
	From    string
	To      string
	Subject string
	Body    string
}

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}

Servus von {{.From}}
`

const LOG_TAG = "MAIL_CLIENT"

var (
	emailUser EmailUser
	smtpAuth  smtp.Auth
)

func Init() {
	if smtpAuth != nil {
		log.Printf("%v Mail was already setup.", LOG_TAG)
		return
	}

	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	url := os.Getenv("MAIL_URL")
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))

	if err != nil {
		log.Panicf("%v MAIL_PORT must be a number", LOG_TAG)
		os.Exit(1)
	}

	if username == "" || password == "" || url == "" || port == 0{
		log.Panicf("%v No Mail-Configuration found! Exiting.", LOG_TAG)
		os.Exit(1)
	}

	emailUser.Username = username
	emailUser.Password = password
	emailUser.EmailServer = url
	emailUser.Port = port

	smtpAuth = smtp.PlainAuth("",
		emailUser.Username,
		emailUser.Password,
		emailUser.EmailServer)

	log.Printf("%v Setup Mail-Client for User %v on Server %v", LOG_TAG, emailUser.Username, emailUser.EmailServer)
}

func SendMail(topic string, message string, receivers []string) error{
	if smtpAuth == nil {
		log.Panicf("%v Mail-Cleitn was not setup correctly! Exiting.", LOG_TAG)
		os.Exit(1)
	}

	context := &SmtpTemplateData{
		"Weisswurst-Systems",
		strings.Join(receivers[:],", "),
		topic,
		message}
	t := template.New("emailTemplate")
	t, err := t.Parse(emailTemplate)
	if err != nil {
		log.Fatalf("%v error trying to parse mail template", LOG_TAG)
		return err
	}

	var doc bytes.Buffer
	err = t.Execute(&doc, context)
	if err != nil {
		log.Fatalf("%v error trying to execute mail template", LOG_TAG)
		return err
	}

	err = smtp.SendMail(emailUser.EmailServer+":"+strconv.Itoa(emailUser.Port),
		smtpAuth,
		emailUser.Username,
		receivers,
		doc.Bytes())
	if err != nil {
		log.Fatalf("%v Error sending mail: %v ", LOG_TAG, err)
		return err
	}
	return nil
}