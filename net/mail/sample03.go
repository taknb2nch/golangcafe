package main

import (
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/mail"
	"os"
	"strings"
)

func main() {
	// メールの内容をそのまま保存したmail2.emlを同じディレクトリに用意してください。
	var f *os.File
	var err error

	if f, err = os.Open("./mail2.eml"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	defer f.Close()

	var message *mail.Message

	if message, err = mail.ReadMessage(f); err != nil {
		fmt.Fprintf(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	contentType := message.Header.Get("content-type")

	parseBody(message.Body, contentType)
}

func parseBody(body io.Reader, contentType string) {
	var mediatype string
	var params map[string]string
	var err error

	if mediatype, params, err = mime.ParseMediaType(contentType); err != nil {
		fmt.Fprintf(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	switch strings.Split(mediatype, "/")[0] {
	case "text", "html":
		fmt.Printf("%v charset=%v\n\n", "text or html.", params["charset"])

	case "message":
		fmt.Printf("message: do something \n\n")

	case "multipart":
		//fmt.Printf("%v\n\n", "multipart")
		boundary := params["boundary"]

		r := multipart.NewReader(body, boundary)

		var part *multipart.Part

		for {
			if part, _ = r.NextPart(); part == nil {
				// err では判定できません。
				break
			}

			contentType := part.Header.Get("content-type")

			parseBody(part, contentType)

			part.Close()
		}

	default:
		// 画像などの添付ファイルなど。ファイルに書き出します。
		fmt.Printf("%v, filename=%v\n\n", mediatype, params["name"])
	}
}
