package bigint

import (
	"app/models"
	"errors"
	"strings"
)

func NewInt(num string) (models.BigInt, error) {
	err := false
	allowed := "1234567890"
	if strings.HasPrefix(num, "-") {
		num = strings.Replace(num, "-", "", 1)
	}
	if num[0] == '0' {
		err = true
	} 
	arr := strings.Split(num, "")
	for _, val := range arr {
		if !strings.Contains(allowed, val) {
			err = true
			break
		}
	}
	if err {
		return models.BigInt{}, errors.New("Invalid input")
	}
	return models.BigInt{Value: num}, nil
}