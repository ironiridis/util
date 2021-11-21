package randid

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func AlNum(n uint) (string, error) {
	const opaqueIdChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var out strings.Builder
	randmax := big.NewInt(int64(len(opaqueIdChars)))
	for n > 0 {
		randval, err := rand.Int(rand.Reader, randmax)
		if err != nil {
			return "", fmt.Errorf("cannot generate an alphanumeric id: %w", err)
		}
		out.WriteByte(opaqueIdChars[randval.Uint64()])
		n--
	}
	return out.String(), nil
}
