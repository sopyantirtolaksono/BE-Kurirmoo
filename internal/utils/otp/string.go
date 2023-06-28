package otp

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func generateStringNumber(maxDigit int) string {
	if maxDigit < 1 {
		return ""
	}

	rand.Seed(time.Now().UnixNano())

	result := make([]string, maxDigit)
	for i := 0; i < maxDigit; i++ {
		randomNumbers := rand.Intn(9)
		result[i] = fmt.Sprintf("%d", randomNumbers)
	}

	return strings.Join(result, "")
}

func GenerateOtpAndOtpMessage(otpType string) (otpCode, message string) {
	otpCode = generateStringNumber(5)
	message = fmt.Sprintf(OTP_MESSAGE_BASE, otpCode, *OtpTypesView[otpType])
	return
}

func PhoneNumberFormat62(phoneNumber string) string {
	return fmt.Sprintf("62%s", phoneNumber[1:])
}
