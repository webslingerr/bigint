package main

import (
	"app/bigint"
	"fmt"
)

func main() {
	a, err := bigint.NewInt("4354556")
	if err != nil {
		panic(err)
	}

	b, err := bigint.NewInt("87")
	if err != nil {
		panic(err)
	}

	// sum := bigint.Add(a, b)
	// fmt.Println(sum.Value)

	div := bigint.Divide(a, b)
	fmt.Println(div.Value)
}
