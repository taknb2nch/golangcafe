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

	const template_text = `初めてのテンプレート。名前は {{.Name}} です。
`

	tpl := template.Must(template.New("mytemplate").Parse(template_text))

	member := Member{1, "ほげほげ", "Go"}

	if err := tpl.Execute(os.Stdout, member); err != nil {
		fmt.Println(err)
	}

}
