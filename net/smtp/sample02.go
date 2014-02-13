package main

import (
	"fmt"
	"io"
	"net/mail"
	"net/smtp"
	"os"
)

func main() {
	smtpServer := "smtp.example.com:587"
	user := "ユーザ名"
	pass := "パスワード"
	from := mail.Address{"From表示名", "from@example.com"}
	to := []mail.Address{
		mail.Address{"to1表示名", "to1@example.com"},
		mail.Address{"to2表示名", "to2@example.com"}}

	auth := smtp.PlainAuth("", user, pass, "smtp.example.com")

	var client *smtp.Client
	var err error

	if client, err = smtp.Dial(smtpServer); err != nil {
		fatal("Error: %v\n", err)
	}

	defer client.Close()

	if err = client.Hello("localhost"); err != nil {
		fatal("Error: %v\n", err)
	}

	if err = client.Auth(auth); err != nil {
		fatal("Error: %v\n", err)
	}

	for _, addr := range to {
		if err = client.Reset(); err != nil {
			fatal("Error: %v\n", err)
		}

		if err = client.Mail(from.Address); err != nil {
			fatal("Error: %v\n", err)
		}

		if err = client.Rcpt(addr.Address); err != nil {
			fatal("Error: %v\n", err)
		}

		msg := "" +
			"From:" + from.String() + "\r\n" +
			"To:" + addr.String() + "\r\n" +
			"Subject:SMTP Test\r\n" +
			"\r\n" +
			"This is a test mail."

		var w io.WriteCloser

		if w, err = client.Data(); err != nil {
			fatal("Error: %v\n", err)
		}

		if _, err = w.Write([]byte(msg)); err != nil {
			fatal("Error: %v\n", err)
		}

		w.Close()
	}

	if err = client.Quit(); err != nil {
		fatal("Error: %v\n", err)
	}
}

func fatal(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args)
	os.Exit(1)
}