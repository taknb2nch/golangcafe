package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	f1(arr)
}

func f1(arr [10]int) {
	fmt.Println(arr)
}
