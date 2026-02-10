package sendto

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
	"text/template"

	"github.com/TaKieuLong/golang_fresher/global"
	"go.uber.org/zap"
)
const (
	SMTP_HOST = "takieulong@gmail.com"
	SMTP_PORT = 587
	SMTP_USER = "takieulong@gmail.com"
	SMTP_PASSWORD = "bmcn muut aazw ffgv"
)
type EmailAddress struct {
	Name string `json:"name"`
	Address string `json:"address"`
}

type Mail struct{
	From EmailAddress `json:"from"`
	To []string `json:"to"`
	Subject string `json:"subject"`
	Body string `json:"body"`
}
func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\"\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("Body: %s\r\n", mail.Body)
	return msg
}
func SendTextEmailOtp(to []string, from string, otp string) (string, error) {
	contentEmail :=Mail{
		From: EmailAddress{
			Name: "Golang Fresher",
			Address: from,
		},
		To: to,
		Subject: "OTP",
		Body: fmt.Sprintf("Your OTP is %s", otp),
	}
	messageEmail :=BuildMessage(contentEmail)

	auth:= smtp.PlainAuth("", SMTP_USER, SMTP_PASSWORD, SMTP_HOST)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", SMTP_HOST, SMTP_PORT), auth, from, to, []byte(messageEmail))
	if err != nil {
		global.Logger.Error("Send email failed::", zap.Error(err))
		return "", err
	}
	return "", nil
}	


func SendTemplateEmailOtp(
	to []string,
	from string,
	nameTemplate string,
	dataTemplate map[string]interface{},
) error {
	htmlBody, err := getMailTemplate(nameTemplate, dataTemplate)
	if err != nil {
		return err
	}

	return send(to, from, htmlBody)
}

func getMailTemplate(
	nameTemplate string,
	dataTemplate map[string]interface{},
) (string, error) {
	htmlTemplate := new(bytes.Buffer)

	t := template.Must(
		template.New(nameTemplate).
			ParseFiles("templates-email/" + nameTemplate),
	)

	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}

	return htmlTemplate.String(), nil
}

func send(to []string, from string, htmlTemplate string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "test"},
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlTemplate,
	}

	messageMail := BuildMessage(contentEmail)

	// send smtp
	auth := smtp.PlainAuth(
		"",
		SMTP_USER,
		SMTP_PASSWORD,
		SMTP_HOST,
	)

	err := smtp.SendMail(
		SMTP_HOST+":587",
		auth,
		from,
		to,
		[]byte(messageMail),
	)
	if err != nil {
		global.Logger.Error("Email send failed::", zap.Error(err))
		return err
	}

	return nil
}
