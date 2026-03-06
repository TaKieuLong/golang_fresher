package sendto

import "encoding/json"

type MailRequest struct {
	To         string `json:"to"`
	Subject    string `json:"subject"`
	Body       string `json:"body"`
	Attachment string `json:"attachment"`
}

func SendMailToJavaByAPI(otp string, email string, purpose string) error {
	postURL := "http://localhost:8080/email/send_text"

	//Data Json
	mailRequest := MailRequest{
		To:         email,
		Subject:    "VerifyOTP" + purpose,
		Body:       otp,
		Attachment: "path/to/email",
	}

	//convert struct to json
	jsonMailRequestBody, err := json.Marshal(mailRequest)
	if err != nil {
		return err
	}
	return nil
}
