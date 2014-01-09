package main

import "fmt"

func main() {
	//
	// 1.1.2でもこの方法は可能だったようです。
	//v := 1.0
	//fmt.Printf("%T\n", v)

	//
	v1, v2, v3, v4, v5 := "あ", "い", "う", "え", "お"

	fmt.Printf("%v, %v, %v, %v, %v\n", v1, v2, v3, v4, v5)

	fmt.Printf("%[5]v, %[5]v, %[3]v, %[3]v, %[1]v\n", v1, v2, v3, v4, v5)

	fmt.Printf("%[2]v, %v, %[1]v, %v, %v\n", v1, v2, v3, v4, v5)
}
