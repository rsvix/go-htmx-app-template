package hash

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
)

func GenerateActivationToken() (string, error) {
	h := sha1.New()
	h.Write([]byte("12"))
	idHash := hex.EncodeToString(h.Sum(nil))
	log.Println(idHash)

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
