package auth

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

// SAMPLE PARAMS
// p := &params{
// 	memory:      64 * 1024,
// 	iterations:  3,
// 	parallelism: 2,
// 	saltLength:  16,
// 	keyLength:   32,
// }

func generateHashFromPassword(password string, p *params) (hash []byte, err error) {
	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return nil, err
	}

	hash = argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	return hash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
