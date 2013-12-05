package main

import (
	"fmt"
	"os"
	"text/template"
)

type Member struct {
	Id     int
	Name   string
	Groups []Group
}

type Group struct {
	Code   string
	Name   string
	Leader bool
}

func (g Group) Display() string {
	return fmt.Sprintf("***%s***", g.Name)
}

func main() {

	tpl := template.Must(template.ParseFiles("sample03.tpl"))

	member := Member{1, "ほげほげ",
		[]Group{
			Group{"01", "GDGChugoku", true},
			Group{"02", "OITEC", false},
		},
	}

	if err := tpl.Execute(os.Stdout, member); err != nil {
		fmt.Println(err)
	}

}
