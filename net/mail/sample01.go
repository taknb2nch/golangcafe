package main

import (
	"fmt"
	"net/mail"
)

func main() {
	parse("Barry Gibbs <bg@example.com>")
	parse("\"Example.co.jp\" <ship-confirm@example.co.jp>")
	parse("hoge..hoge@example.ne.jp")
	parse("hoge..@example.ne.jp")
	parse("..hoge@example.ne.jp")
	parse("")
	parse("account.example.com")
	parse("=?utf-8?q?B=C3=B6b?= <bob@example.com>")
}

func parse(address string) {
	if addr, err := mail.ParseAddress(address); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%v -> Name: %v, Address: %v\n", address, addr.Name, addr.Address)
	}
}
