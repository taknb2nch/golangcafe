package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	fa(array)

	fmt.Printf("array: %v\n", array)

	slice := array[:]

	fs(slice)

	fmt.Printf("array: %v\n", array)
	fmt.Printf("slice: %v\n", slice)
}

func fa(arr [5]int) {
	for i, _ := range arr {
		arr[i] *= 10
	}
}

func fs(slc []int) {
	for i, _ := range slc {
		slc[i] *= 1000
	}
}
