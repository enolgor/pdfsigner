package db

import (
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/scrypt"
)

const (
	scryptN = 1 << 15 // 32768
	scryptR = 8
	scryptP = 2
	keyLen  = 32
	saltLen = 16
)

func Derive(password string, saltstr string) (key []byte, salt []byte, err error) {
	if password == "" {
		return nil, nil, nil
	}
	if saltstr != "" {
		if salt, err = hex.DecodeString(saltstr); err != nil {
			return
		}
	} else {
		if salt, err = randomBytes(saltLen); err != nil {
			return
		}
	}
	key, err = scrypt.Key([]byte(password), salt, scryptN, scryptR, scryptP, keyLen)
	return
}

func randomBytes(length int) ([]byte, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}
