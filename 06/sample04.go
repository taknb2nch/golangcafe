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

	tpl := template.Must(template.ParseFiles("sample04_1.tpl", "sample04_2.tpl"))

	tpl1 := tpl.Lookup("sample04_2.tpl")

	member := Member{1, "ほげほげ", "Go"}

	if err := tpl1.Execute(os.Stdout, member); err != nil {
		fmt.Println(err)
	}

	if err := tpl.ExecuteTemplate(os.Stdout, "sample04_2.tpl", member); err != nil {
		fmt.Println(err)
	}	

}
