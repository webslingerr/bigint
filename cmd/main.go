package main

import (
	"app/bigint"
	"fmt"
)

func main() {
	a, err := bigint.NewInt("12")
	if err != nil {
		panic(err)
	}

	b, err := bigint.NewInt("24")
	if err != nil {
		panic(err)
	}

	sum := bigint.Add(a, b)
	fmt.Println(sum.Value)

	sub := bigint.Sub(a, b)
	fmt.Println(sub.Value)
}