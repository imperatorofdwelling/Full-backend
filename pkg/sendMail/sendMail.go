package sendMail

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
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

	subject := "Subject: " + title
	to := []string{userMail}
	body := fmt.Sprintf("Your code: %s", userOTP)

	message := []byte(subject + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + "Someone Of Dwellers" + "\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n" +
		"\r\n" +
		body)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", smtpHost, smtpPort),
		auth,
		smtpUser,
		to,
		message,
	)

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
