package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateVerificationCode генерирует 6-значный код
func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
