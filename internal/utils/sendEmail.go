package utils

import (
	"fmt"
	"net/smtp"
	config "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/config"
)

func SendEmail(recipient string, msg string) (string, error) {
	config, _ := config.LoadConfig(".")
	auth := smtp.PlainAuth(
		"",
		config.Email,
		config.EmailPassword,
		"smtp.gmail.com",
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		config.Email,
		[]string{recipient},
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}

	return "", nil
}
