package email

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"
)

func SendActivationMail(to string, uidHash string) (int, error) { // оптравка письма активации на почту пользователя
	var smtpEmail string = os.Getenv("SMTP_EMAIL")
	var smtpPassword string = os.Getenv("SMTP_PASSWORD")

	host := "mail.inbox.lv"
	port := "587"
	address := host + ":" + port
	auth := smtp.PlainAuth("", smtpEmail, smtpPassword, host)
	subject := "Email Verificaion"
	msg := []byte(
		"From: " + smtpEmail + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject +
			"\r\n" + "MIME: MIME-version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
			`<html> 
			<h1>Нажмите ссылку, чтобы подтвердить адрес электронной почты</h1>
			<a href="` + "http://localhost:9000/#" + `/verify/verify-email/` + uidHash + `">подтвердить адрес электронной почты</a>
		</html>`,
	)
	err := smtp.SendMail(address, auth, smtpEmail, []string{to}, msg)
	if err != nil {
		return http.StatusOK, fmt.Errorf("ошибка отправки письма активации на вашу почту")
	}

	return http.StatusOK, nil
}

func SendRecoveryPasswordMail(to string, uidHash string) (int, error) { // оптравка письма для восстановления пароля на почту пользователя
	var smtpEmail string = os.Getenv("SMTP_EMAIL")
	var smtpPassword string = os.Getenv("SMTP_PASSWORD")

	host := "mail.inbox.lv"
	port := "587"
	address := host + ":" + port
	auth := smtp.PlainAuth("", smtpEmail, smtpPassword, host)
	subject := "Password Recovery"
	msg := []byte(
		"From: " + smtpEmail + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject +
			"\r\n" + "MIME: MIME-version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
			`<html> 
			<h1>Нажмите ссылку для восстановления пароля</h1>
			<a href="` + "http://localhost:9000/#" + `/recovery/recovery-password/` + uidHash + `">восстановить пароль</a>
		</html>`,
	)

	err := smtp.SendMail(address, auth, smtpEmail, []string{to}, msg)
	if err != nil {
		return http.StatusOK, fmt.Errorf("ошибка отправки письма восстановления пароля на вашу почту")
	}

	return http.StatusOK, nil
}
