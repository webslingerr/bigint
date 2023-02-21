package bigint

import "app/models"

func Abs(x *models.BigInt) models.BigInt {
	if x.Value[0] == '-' {
		return models.BigInt{
			Value: x.Value[1:],
		}
	}
	if x.Value[0] == '+' {
		return models.BigInt{
			Value: x.Value[1:],
		}
	}
	return models.BigInt{Value: x.Value}
}