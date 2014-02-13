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
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	defer client.Close()

	if err = client.Hello("localhost"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err = client.Auth(auth); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for _, addr := range to {
		if err = client.Reset(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if err = client.Mail(from.Address); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if err = client.Rcpt(addr.Address); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		msg := "" +
			"From:" + from.String() + "\r\n" +
			"To:" + addr.String() + "\r\n" +
			"Subject:SMTP Test\r\n" +
			"\r\n" +
			"This is a test mail."

		var w io.WriteCloser

		if w, err = client.Data(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if _, err = w.Write([]byte(msg)); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		w.Close()
	}

	if err = client.Quit(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
