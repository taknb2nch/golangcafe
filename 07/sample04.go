package main

import "fmt"

func main() {
	m1()
	fmt.Println()
	m2()
	fmt.Println()
	m3()
}

func m1() {
	arr := setup()

	slice := arr[2:4]

	appendAndOutput(slice)

	fmt.Println(arr)
}

func m2() {
	arr := setup()

	slice := arr[2:4:7]

	appendAndOutput(slice)

	fmt.Println(arr)
}

func m3() {
	arr := setup()

	slice := arr[2:4:4]

	appendAndOutput(slice)

	fmt.Println(arr)
}

func setup() [10]int {
	return [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func appendAndOutput(slice []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	for i := 0; i < 10; i++ {
	slice = append(slice, 100)
	}

	fmt.Println(slice)
}