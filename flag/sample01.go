package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	x := flag.Bool("x", false, "xフラグです。true または false")
	y := flag.Bool("y", false, "yフラグです。true または false")

	var z int
	flag.IntVar(&z, "z", 100, "整数を指定します。")

	var g string
	flag.StringVar(&g, "g", "empty", "文字列を入力します。")
	flag.StringVar(&g, "gopher", "empty", "文字列を入力します。")

	flag.Parse()

	fmt.Printf("x: %v, y: %v, z: %v\n", *x, *y, z)
	fmt.Printf("args: %v\n", flag.Args())

	//
	var args = make([]string, 0)

	flag.Visit(func(f *flag.Flag) {
		args = append(args, fmt.Sprintf("%s: %v", f.Name, f.Value))
	})

	fmt.Println("visit: [" + strings.Join(args, ", ") + "]")

	//
	args = make([]string, 0)

	flag.VisitAll(func(f *flag.Flag) {
		args = append(args, fmt.Sprintf("%s: %v", f.Name, f.Value))
	})

	fmt.Println("visitall: [" + strings.Join(args, ", ") + "]")
}
