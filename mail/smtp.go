package mail

import (
	"log"
	"net/smtp"
	"os"
	"strconv"
)

type smtpService struct {
	user emailUser
	auth smtp.Auth
}

type emailUser struct {
	username   string
	password   string
	smtpServer string
	port       int
}

func NewSMTPService() Service {
	var service smtpService
	port, err := strconv.Atoi(os.Getenv("mail.port"))
	if err != nil {
		log.Printf("%v mail.port must be a number", LOG_TAG)
	}
	service.user = emailUser{
		username:   os.Getenv("mail.username"),
		password:   os.Getenv("mail.password"),
		smtpServer: os.Getenv("mail.smtpServer"),
		port:       port,
	}
	service.auth = smtp.PlainAuth("",
		service.user.username,
		service.user.password,
		service.user.smtpServer)
	log.Printf("%v Setup Mail-Client for User %v on Server %v", LOG_TAG, service.user.username, service.user.username)

	return &service
}

func (service *smtpService) Send(mail Mail) error {
	address := service.user.smtpServer + ":" + strconv.Itoa(service.user.port)
	err := smtp.SendMail(address,
		service.auth,
		"Weisswurst Systems",
		mail.Receivers,
		mail.Content)
	if err != nil {
		log.Printf("%v Error sending mail: %v ", LOG_TAG, err)
		return err
	}
	return nil
}
