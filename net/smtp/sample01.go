package main

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"os"
)

func main() {
	// 適宜変更してください
	smtpServer := "smtp.example.com:587"
	user := "ユーザ名"
	pass := "パスワード"
	from := mail.Address{"From表示名", "from@example.com"}
	to := mail.Address{"to表示名", "to@example.com"}
	cc := mail.Address{"cc表示名", "cc@example.com"}
	receivers := []string{to.Address, cc.Address}

	msg := "" +
		"From:" + from.String() + "\r\n" +
		"To:" + to.String() + "\r\n" +
		"Cc:" + cc.String() + "\r\n" +
		"Subject:SMTP Test\r\n" +
		"\r\n" +
		"This is a test mail."

	auth := smtp.PlainAuth("", user, pass, "smtp.example.com")
	//auth := smtp.CRAMMD5Auth(user, pass)

	err := smtp.SendMail(smtpServer, auth, from.Address, receivers, []byte(msg))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
