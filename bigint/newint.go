package bigint

import (
	"app/models"
)

func NewInt(num string) (models.BigInt, error) {
	return models.BigInt{Value: num}, nil
}