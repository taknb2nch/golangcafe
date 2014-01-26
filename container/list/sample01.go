package main

import (
	"fmt"
	"container/list"
)

func main() {
	l := list.New()

	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}

	fmt.Printf("要素数: %d\n", l.Len())

	//
	e := l.Front()

	fmt.Printf("先頭 value: %v, prev: %v, next: %v\n", e.Value, e.Prev(), e.Next())
	fmt.Printf("要素数: %d\n", l.Len())

	//
	e = l.Back()

	fmt.Printf("最後 value: %v, prev: %v, next: %v\n", e.Value, e.Prev(), e.Next())
	fmt.Printf("要素数: %d\n", l.Len())

	//
	fmt.Println("リスト要素を列挙")
	for v := l.Front(); v != nil; v = v.Next() {
		fmt.Printf("%v ", v.Value)
	}
	fmt.Println()

	//
	fmt.Println("先頭を削除")
	e = l.Front()
	l.Remove(e)

	fmt.Printf("要素数: %d\n", l.Len())

	//
	fmt.Println("リスト要素を列挙")
	for v := l.Front(); v != nil; v = v.Next() {
		fmt.Printf("%v ", v.Value)
	}
	fmt.Println()
}

