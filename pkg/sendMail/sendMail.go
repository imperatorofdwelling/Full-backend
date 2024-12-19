package sendMail

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
)

func SimpleEmailSend(userMail, userOTP string) error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	// Формируем абсолютный путь к .env.mail
	envFilePath := filepath.Join(projectRoot, ".env.mail")

	err = godotenv.Load(envFilePath)
	if err != nil {
		return fmt.Errorf("Error loading .env file: %w", err)
	}

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

	subject := "Subject: Registration"
	to := []string{userMail}
	body := fmt.Sprintf("Your code: %s", userOTP)

	message := []byte(subject + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + smtpUser + "\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n" +
		"\r\n" +
		body)

	err = smtp.SendMail(
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
