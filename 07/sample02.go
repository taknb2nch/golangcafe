package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int32 = 1

	// reflectパッケージを使った方法
	typeName1 := reflect.ValueOf(i).Type().Name()
	// 1.2での追加
	typeName2 := fmt.Sprintf("%T", i)

	fmt.Println(typeName1, typeName2)
}
