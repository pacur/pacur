package utils

import (
	"crypto/rand"
	"math/big"
	mathrand "math/rand"
	"time"
)

func seedRand() {
	n, err := rand.Int(rand.Reader, big.NewInt(9223372036854775806))
	if err != nil {
		mathrand.Seed(time.Now().UTC().UnixNano())
		return
	}

	mathrand.Seed(n.Int64())
	return
}

func init() {
	seedRand()
}
