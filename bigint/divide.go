package bigint

import (
	"app/models"
)

func Divide(a, b models.BigInt) models.BigInt {
	val1, val2 := a.Value, b.Value
	if isSmaller(val1, val2) {
		return models.BigInt{Value: "0"}
	}

	c := models.BigInt{}
	for {
		res := Sub(a, b)
		c = Add(c, models.BigInt{Value: "1"})
		if res.Value[0] == '0' {
			res.Value = res.Value[1:]
		}
		if isSmaller(res.Value, b.Value) {
			return models.BigInt{Value: c.Value}
		}
		a = res
	}
}
