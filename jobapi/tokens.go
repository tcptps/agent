package jobapi

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateToken(length int) (string, error) {
	lettersLen := len(letters)
	bigLettersLen := big.NewInt(int64(lettersLen))
	b := make([]byte, length)
	for i := range b {
		bigRand, err := rand.Int(rand.Reader, bigLettersLen)
		if err != nil {
			return "", fmt.Errorf("generating crypto-random big.Int: %w", err)
		}
		b[i] = letters[bigRand.Int64()%int64(lettersLen)] // mod letterslen just in case rand.Int returns a number larger than lettersLen, don't want to have a panic
	}
	return string(b), nil
}
