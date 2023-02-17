package bigint

import (
	"app/models"
	"strconv"
)

func isSmaller(a, b string) bool {
	n1, n2 := len(a), len(b)
	if n1 < n2 {
		return true
	} else if n1 > n2{
		return false
	}
	for i:=0; i<n1; i++ {
		if (a[i] < b[i]) {
			return true
		} else if (a[i] > b[i]) {
			return false
		}
	}
	return false
}

func Sub(a, b models.BigInt) models.BigInt {
	val1, val2 := a.Value, b.Value
	if isSmaller(val1, val2) {
		val1, val2 = val2, val1
	}

	rVal1, rVal2 := reverse(val1), reverse(val2)
	n1, n2 := len(val1), len(val2)

	answer := ""

	carry := 0

	for i:=0; i<n2; i++ {
		num1, _ := strconv.Atoi(string(rVal1[i]))
		num2, _ := strconv.Atoi(string(rVal2[i]))
		sub := num1-num2-carry
		if sub < 0 {
			sub += 10
			carry = 1
		} else {
			carry = 0
		}
		answer += strconv.Itoa(sub)
	}

	for i:=n2; i<n1; i++ {
		num1, _  := strconv.Atoi(string(rVal1[i]))
		sub := num1-carry
		if sub < 0 {
			sub += 10
			carry = 1
		} else {
			carry = 0
		}
		answer += strconv.Itoa(sub)
	}
	rAnswer := reverse(answer)
	return models.BigInt{
		Value: rAnswer,
	}
}
