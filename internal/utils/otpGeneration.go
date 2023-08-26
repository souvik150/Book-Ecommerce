package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateOTP(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("invalid OTP length")
	}

	const digits = "0123456789"
	max := big.NewInt(int64(len(digits)))

	otp := make([]byte, length)
	for i := range otp {
		randomIndex, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		otp[i] = digits[randomIndex.Int64()]
	}

	return string(otp), nil
}
