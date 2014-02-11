package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	x := flag.Bool("x", false, "xフラグです。true または false")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "パラメータ解析中にエラーが発生しました。\n")
		flag.PrintDefaults()
		os.Exit(2)
	}

	flag.Parse()

	fmt.Printf("x: %v\n", *x)
}
