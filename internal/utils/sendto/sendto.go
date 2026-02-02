package sendto

import "github.com/TaKieuLong/golang_fresher/internal/model"

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
	msg += fmt.Sprintf("To: %s\r\n", mail.To)
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("Body: %s\r\n", mail.Body)
	return msg
}
func SendTextEmailOtp(to []string, from string, otp string) error {
	contentEmail :=Mail{
		From: EmailAddress{
			Name: "Golang Fresher",
			Address: from,
		},
		To: to,
		Subject: "OTP",
		Body: fmt.Sprintf("Your OTP is %s", otp),
	}
	messageEmail :=
}