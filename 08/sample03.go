package main

import "fmt"

func main() {
	v1, v2 := 1, "2"

	f1(v1)
	f1(v2)
}

func f1(v string) {
	fmt.Println(v)
}

func f1(v int) {
	fmt.Println(v)
}
