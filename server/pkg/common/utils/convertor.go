package utils

import (
	"strconv"
)

func ConvertStringToUint(str string) (uint64, error) {
	convID, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return convID, nil
}

func ConvertStringToBool(boolean string) (bool, error) {
	result, err := strconv.ParseBool(boolean)
	if err != nil {
		return false, err
	}

	return result, nil
}
