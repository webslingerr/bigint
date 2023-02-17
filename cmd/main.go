package main

import (
	"app/bigint"
	"fmt"
)

func main() {
	a, err := bigint.NewInt("988847123412385995937737458959")
	if err != nil {
		panic(err)
	}

	b, err := bigint.NewInt("21231231231231231231231231233")
	if err != nil {
		panic(err)
	}

	sum := bigint.Add(a, b)
	fmt.Println(sum.Value)
}