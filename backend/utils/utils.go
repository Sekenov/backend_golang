package utils

import (
	"fmt"
	"net/smtp"
)

// SendVerificationCode отправляет код на указанный email
func SendVerificationCode(toEmail, code string) error {
	from := "temohatemoha25@gmail.com" // Ваш email-адрес
	password := "pfcn kwlj gmxj dwbw" // Пароль приложения Gmail

	// Настройки SMTP
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Текст сообщения
	message := []byte(fmt.Sprintf("Subject: Verification Code\n\nYour verification code is: %s", code))

	// Авторизация
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Отправка сообщения
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, message)
	if err != nil {
		// Улучшенная обработка ошибок с дополнительной информацией
		return fmt.Errorf("failed to send email to %s: %v", toEmail, err)
	}

	fmt.Println("Отправляем код подтверждения:", code, "на", toEmail)
	fmt.Println("Используем SMTP-хост:", smtpHost, "и порт:", smtpPort)
	
	return nil
}
