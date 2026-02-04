package sendto

import (
	"fmt"
	"net/smtp"
	"strings"

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