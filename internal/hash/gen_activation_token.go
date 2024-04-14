package hash

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"
)

func GenerateActivationToken() (string, error) {
	bigInt, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return "", err
	}

	tenDigitNum := bigInt.Int64() + 1000000000
	tenDigitStr := fmt.Sprintf("%06d", tenDigitNum)

	hash := sha256.Sum256([]byte(tenDigitStr))
	// hash := sha512.Sum512([]byte(tenDigitStr))
	activation_token := strings.TrimSpace(fmt.Sprintf("%x\n", hash))

	return activation_token, nil
}
