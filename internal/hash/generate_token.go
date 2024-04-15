package hash

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
)

func GenerateToken(forActivation bool, id string) (string, error) {
	var src []byte
	if forActivation {
		src = []byte("activate@" + id)
	} else {
		src = []byte("resetpwd@" + id)
	}
	idEncoded := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(idEncoded, src)

	bigInt, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return "", err
	}
	tenDigitNum := bigInt.Int64() + 1000000000
	tenDigitStr := fmt.Sprintf("%06d", tenDigitNum)
	// SHA 256
	// hash := sha256.Sum256([]byte(tenDigitStr))
	// SHA 512
	hash := sha512.Sum512([]byte(tenDigitStr))
	EncodedRandomNumberHash := hex.EncodeToString(hash[:])

	token := strings.TrimSpace(fmt.Sprintf("%sO%s", EncodedRandomNumberHash, idEncoded))
	log.Printf("Encoded: %v", token)
	return token, nil
}

func DecodeToken(s string) (string, error) {
	decoded, err := hex.DecodeString(strings.Split(s, "O")[1])
	if err != nil {
		log.Println(err)
		return "", err
	}
	decodedStr := string(decoded[:])
	mode := strings.Split(decodedStr, "@")[0]
	id := strings.Split(decodedStr, "@")[1]

	log.Printf("Decoded: %v", decodedStr)
	log.Printf("Mode: %v", mode)
	log.Printf("Id: %v", id)

	return id, nil
}
