package shared

import (
	"fmt"
	"strings"
)

func ValidateDigitFloatingPoint(n float64, digit int) bool {
	splitedString := strings.Split(fmt.Sprintf("%v", n), ".")

	if len(splitedString) > 1 {
		return len(splitedString[1]) > digit
	}

	return false
}
