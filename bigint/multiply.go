package bigint

import (
	"app/models"
	"strconv"
)

func max(a, b int) int {
	if a<b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func Multiply(a, b models.BigInt) models.BigInt {
	val1 := a.Value
	val2 := b.Value
	if val1 == "0" || val2 == "0" {
		return models.BigInt{Value: "0"}
	}

	m, n, carry := len(val1)-1, len(val2)-1, 0
	answer := ""

	for i:=0; i<=m+n || carry != 0; i++ {
		for j:=max(0, i-n); j<=min(i, m); j++ {
			num1, _ := strconv.Atoi(string(val1[m-j]))
			num2, _ := strconv.Atoi(string(val2[n-i+j]))
			carry += num1 * num2
		}
		answer += strconv.Itoa(carry % 10)
		carry /= 10
	}
	rAnswer := reverse(answer)
	return models.BigInt{Value: rAnswer}
}