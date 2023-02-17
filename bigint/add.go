package bigint

import (
	"app/models"
	"strconv"
)

func Add(a, b models.BigInt) models.BigInt {
	val1 := a.Value
	val2 := b.Value
	if len(val1) > len(val2) {
		val1, val2 = val2, val1
	}
	n1, n2 := len(val1), len(val2)

	rVal1, rVal2 := reverse(val1), reverse(val2)

	carry := 0

	answer := ""

	for i:=0; i<n1; i++ {
		num1, _ := strconv.Atoi(string(rVal1[i]))
		num2, _ := strconv.Atoi(string(rVal2[i]))
		sum := num1+num2+carry
		answer += strconv.Itoa(sum%10)
		carry = sum/10
	}

	for i:=n1; i<n2; i++ {
		num2, _ := strconv.Atoi(string(rVal2[i]))
		sum := num2+carry
		answer += strconv.Itoa(sum%10)
		carry = sum/10
	}

	if carry != 0 {
		answer += strconv.Itoa(carry)
	}

	rAnswer := reverse(answer)
	return models.BigInt{Value: rAnswer}
}