package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/go-clean-arch-boilerplate/models"
	"gopkg.in/gomail.v2"
)

//SendMail func
func SendMail(payload interface{}, mailType string) {
	var err error

	result, toEmail, subject := SetData(payload, mailType)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("MAIL_FROM"))
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", result)

	dialer := gomail.NewDialer(
		os.Getenv("MAIL_HOST"),
		587,
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASS"),
	)

	err = dialer.DialAndSend(mailer)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Mail sent!")
}

//SetData func
func SetData(payload interface{}, mailType string) (string, string, string) {
	var toEmail string
	var subject string

	if mailType == "signup" {
		data := payload.(models.SignupMailData)

		t, err := template.ParseFiles(fmt.Sprintf("library/mail/template/%s", data.Template))
		if err != nil {
			log.Println(err)
		}

		var tpl bytes.Buffer
		if err := t.Execute(&tpl, data); err != nil {
			log.Println(err)
		}

		result := tpl.String()

		toEmail = data.Email
		subject = "Register User PesenYuk!"

		return result, toEmail, subject
	}

	if mailType == "forgot_password" {
		data := payload.(models.ForgotPasswordMailData)

		t, err := template.ParseFiles(fmt.Sprintf("library/mail/template/%s", data.Template))
		if err != nil {
			log.Println(err)
		}

		var tpl bytes.Buffer
		if err := t.Execute(&tpl, data); err != nil {
			log.Println(err)
		}

		result := tpl.String()

		toEmail = data.Email
		subject = "Forgot Password PesenYuk!"

		return result, toEmail, subject
	}

	return "", toEmail, subject
}
