package helpers

import (
	"errors"
	"os"
	"strings"

	"gopkg.in/gomail.v2"
)

var GMAILPASS = ""

func SendMail(title, email, urllink string) error {
	if title == "" || email == "" {
		return errors.New("title and email must be fill")
	}
	const CONFIG_SMTP_HOST = "smtp.gmail.com"
	const CONFIG_SMTP_PORT = 587
	const CONFIG_SENDER_NAME = "Lapak UMKM <findryankpradana@gmail.com>"
	const CONFIG_AUTH_EMAIL = "findryankpradana@gmail.com"
	var CONFIG_AUTH_PASSWORD = GMAILPASS

	content, errContent := os.ReadFile("utils/files/forgetpassword.html")
	if errContent != nil {
		return errContent
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email)
	mailer.SetAddressHeader("Cc", email, title)
	mailer.SetHeader("Subject", "Forget Password")

	contentStr := string(content)
	contentStr = strings.Replace(contentStr, "{{title}}", title, -1)
	contentStr = strings.Replace(contentStr, "{{email}}", email, -1)
	contentStr = strings.Replace(contentStr, "{{urllink}}", urllink, -1)
	mailer.SetBody("text/html", contentStr)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}
