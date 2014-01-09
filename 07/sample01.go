package main

import (
	"fmt"
	"math/big"
)

func main() {
	i, j, k := big.NewInt(1), big.NewInt(2), big.NewInt(0)

	x := i.Add(i, j)
	k.Add(i, j)

	fmt.Println(i, j, k, x)
}
