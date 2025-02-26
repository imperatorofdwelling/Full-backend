package sendMail

import (
	"fmt"
	"net/smtp"
	"os"
)

func SimpleEmailSend(userMail, userOTP, title string) error {
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	if smtpUser == "" || smtpPassword == "" || smtpHost == "" || smtpPort == "" {
		return fmt.Errorf("SMTP configuration is missing in .env file")
	}

	auth := smtp.PlainAuth(
		"",
		smtpUser,
		smtpPassword,
		smtpHost,
	)

	// Создаем сообщение с правильным порядком заголовков
	message := []byte(fmt.Sprintf("From: %s <%s>\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/plain; charset=UTF-8\r\n"+
		"\r\n"+
		"Your code: %s", "Someone Of Dwellers", smtpUser, userMail, title, userOTP))

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", smtpHost, smtpPort),
		auth,
		smtpUser,
		[]string{userMail},
		message,
	)

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
