package main

import (
	"app/bigint"
	"fmt"
)

func main() {
	a, err := bigint.NewInt("2556456")
	if err != nil {
		panic(err)
	}

	b, err := bigint.NewInt("456656")
	if err != nil {
		panic(err)
	}

	// sum := bigint.Add(a, b)
	// fmt.Println(sum.Value)

	// div := bigint.Divide(a, b)
	// fmt.Println(div.Value)
	mod := bigint.Mod(a, b)
	fmt.Println(mod.Value)
}