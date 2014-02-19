package main

import (
	// hg clone https://code.google.com/p/go.text/
	"code.google.com/p/go.text/encoding/japanese"
	"code.google.com/p/go.text/transform"
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
		fatal("Error: %v\n", err)
	}

	defer f.Close()

	var message *mail.Message

	if message, err = mail.ReadMessage(f); err != nil {
		fatal("Error: %v\n", err)
	}

	for k, v := range message.Header {
		fmt.Printf("%v\n%v\n\n", k, v)
	}

	fmt.Printf("%v -> %v\n", "content-type", message.Header.Get("content-type"))

	// addrs, err := message.Header.AddressList("from")

	// fmt.Printf("1, %v\n", err)

	// for _, addr := range addrs {
	// 	fmt.Printf("Name: %v, Address: %v\n", addr.Name, addr.Address)
	// }

	var body []byte

	// if body, err = ioutil.ReadAll(message.Body); err != nil {
	// 	fatal("Error: %v\n", err)
	// }

	// utf8エンコーディングの場合は変換処理は不要です。
	// dst := make([]byte, len(body)*2)
	// var dlen int
	// transformer := japanese.ISO2022JP.NewDecoder()

	// if dlen, _, err = transformer.Transform(dst, body, true); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: ", err)
	// 	os.Exit(1)
	// }

	//fmt.Printf("%v\n", string(dst[:dlen]))

	// transform.NewReaderを使うほうがスマート!
	if body, err = ioutil.ReadAll(
		transform.NewReader(message.Body, japanese.ISO2022JP.NewDecoder())); err != nil {
		fatal("Error: %v\n", err)
	}

	fmt.Printf("%v\n", string(body))
}

func fatal(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args)
	os.Exit(1)
}
