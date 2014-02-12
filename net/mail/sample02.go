package main

import (
	// hg clone https://code.google.com/p/go.text/
	"code.google.com/p/go.text/encoding/japanese"
	"fmt"
	"io/ioutil"
	"net/mail"
	"os"
)

func main() {
	// メールの内容をそのまま保存したmail.emlを同じディレクトリに用意してください。
	var f *os.File
	var err error

	if f, err = os.Open("./mail.eml"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	defer f.Close()

	var message *mail.Message

	if message, err = mail.ReadMessage(f); err != nil {
		fmt.Fprintf(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	for k, v := range message.Header {
		fmt.Printf("%v\n%v\n\n", k, v)
	}

	fmt.Printf("%v -> %v\n", "content-type", message.Header.Get("content-type"))

	addrs, err := message.Header.AddressList("from")

	fmt.Printf("1, %v\n", err)

	for _, addr := range addrs {
		fmt.Printf("Name: %v, Address: %v\n", addr.Name, addr.Address)
	}

	var body []byte

	if body, err = ioutil.ReadAll(message.Body); err != nil {
		fmt.Fprintf(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	// utf8エンコーディングの場合は変換処理は不要です。
	dst := make([]byte, len(body))
	transformer := japanese.ISO2022JP.NewDecoder()

	if _, _, err = transformer.Transform(dst, body, false); err != nil {
		fmt.Fprintf(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n", string(dst))
}
