package otp

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateOTP() string {
	const num = 1000000
	random, _ := rand.Int(rand.Reader, big.NewInt(num))

	return fmt.Sprintf("%06d", random.Int64())
}
