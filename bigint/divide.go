package bigint

import (
	"app/models"
)

func Div(a, b string) string {
	c := models.BigInt{}
	for {
		res := Sub(models.BigInt{Value: a}, models.BigInt{Value: b})
		c = Add(c, models.BigInt{Value: "1"})
		if res.Value[0] == '0' {
			res.Value = res.Value[1:]
		}
		if isSmaller(res.Value, b) {
			return c.Value
		}
		a = res.Value
	}
}

func Divide(a, b models.BigInt) models.BigInt {
	v1, v2 := a.Value, b.Value
	if isSmaller(v1, v2) {
		return models.BigInt{Value: "0"}
	}

	res := ""
	t := models.BigInt{}
	r := models.BigInt{Value: v1}

	for i := 1; ; i++ {
		if !isSmaller(v1[:i], b.Value) {
			res += Div(r.Value, b.Value)
			l := Sub(models.BigInt{Value: v1[:i]}, Multiply(b, models.BigInt{Value: res}))
			if l.Value[0] == '0' {
				t.Value += l.Value[1:]
			} else {
				t.Value += l.Value
			}
			t.Value += v1[i:]
			r.Value = t.Value
			t.Value = ""
			i = 1
			if len(res) == 5 {
				return models.BigInt{Value: res}
			}
		}
	}
}
