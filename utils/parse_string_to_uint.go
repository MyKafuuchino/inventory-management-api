package utils

import (
	"errors"
	"strconv"
)

func ParseStringToUint(stringInt string) (uint, error) {
	id, err := strconv.Atoi(stringInt)
	if err != nil {
		return 0, errors.New(stringInt + " is not a number")
	}
	return uint(id), nil
}
