package main

import (
	"errors"
	"flag"
	"fmt"
)

type myType string

func (m *myType) String() string {
	return string(*m)
}

func (m *myType) Set(value string) error {
	if string(*m) != "" {
		return errors.New("すでに値が設定されています。")
	}

	*m = myType("###_" + value + "_###")

	return nil
}

func main() {
	var m myType
	flag.Var(&m, "mytype", "myTypeのオプションです。")

	flag.Parse()

	fmt.Printf("myType: %v\n", m)
}
