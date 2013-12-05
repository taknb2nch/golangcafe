package main

import (
	"fmt"
	"os"
	"text/template"
)

type Member struct {
	Id   int
	Name string
	Tech string
}

func main() {

	tpl := template.Must(template.ParseFiles("sample05_1.tpl", "sample05_2.tpl"))

	member := Member{1, "ほげほげ", "Go"}

	if err := tpl.Execute(os.Stdout, member); err != nil {
		fmt.Println(err)
	}

}
