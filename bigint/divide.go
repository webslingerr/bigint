package bigint

import (
	"app/models"
	"strconv"
)

func Divide(a, b models.BigInt) models.BigInt {
	answer := ""

	i := 0
	t, _ := strconv.Atoi(string(a.Value[i]))
	divisor, _ := strconv.Atoi(string(b.Value))
	for t < divisor {
		i += 1
		j, _ := strconv.Atoi(string(a.Value[i]))
		t = t * 10 + j
	}

	for len(a.Value) > i {
		answer += strconv.Itoa(t/divisor)
		i++
		j, _ := strconv.Atoi(string(a.Value[i])) 
		t = (t%divisor) * 10 + j
	}

	if len(answer) == 0 {
		return models.BigInt{Value: "0"}
	}
	return models.BigInt{Value: answer}
}