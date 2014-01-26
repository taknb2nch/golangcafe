package main

import (
	"fmt"
	"container/ring"
)

func main() {
	r := ring.New(10)

	for i := 0; i < 10; i++ {
		r.Value = i

		r = r.Next()
	}

	fmt.Printf("要素数: %d\n", r.Len())
	fmt.Printf("現在の値: %v\n", r.Value)

	//
	printElements("すべての要素を列挙", r)

	//
	fmt.Printf("現在の値: %v\n", r.Value)
	fmt.Println("20回Nextしならが要素を列挙")
	for i := 0; i < 20; i++ {
		fmt.Printf("%v ", r.Value)

		r = r.Next()
	}
	fmt.Println()

	//
	fmt.Printf("現在の値: %v\n", r.Value)
	fmt.Println("先頭から5個の要素を削除")

	r1 := r.Unlink(5)

	printElements("残ったすべての要素を列挙", r)
	printElements("削除したすべての要素を列挙", r1)

	//
	fmt.Printf("現在の値: %v\n", r.Value)

	r = r.Link(r1)

	printElements("追加後のすべての要素を列挙", r)

	//
	fmt.Println("同じringをlink")

	r = r.Link(r)

	fmt.Printf("要素数: %d\n", r.Len())
	fmt.Printf("現在の値: %v\n", r.Value)
	printElements("追加後のすべての要素を列挙", r)
}

func printElements(msg string, r *ring.Ring) {
	fmt.Println(msg)
	r.Do(func(v interface{}) {
		fmt.Printf("%v ", v)
		})

	fmt.Println()
}