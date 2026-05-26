package handler

import (
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func SendEmail(email,text string)  {
	message := gomail.NewMessage()

    message.SetHeader("From", "yihunezewdie23@gmail.com")
    message.SetHeader("To", email)
    message.SetHeader("Subject", "Hello from the AchirURL team")

    message.SetBody("text/plain", text)

    dialer := gomail.NewDialer("live.smtp.mailtrap.io", 587, "api", "1a2b3c4d5e6f7g")

    if err := dialer.DialAndSend(message); err != nil {
        fmt.Println("Error:", err)
        panic(err)
    } else {
        fmt.Println("Email sent successfully!")
    }
}