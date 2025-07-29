package store

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/goccy/go-json"
	"golang.org/x/crypto/scrypt"
)

const (
	scryptN = 1 << 15 // 32768
	scryptR = 8
	scryptP = 2
	keyLen  = 32
	saltLen = 16
)

func Derive(password string, saltstr string) (key []byte, newsalt string, err error) {
	var salt []byte
	if password == "" {
		return
	}
	if saltstr != "" {
		if salt, err = hex.DecodeString(saltstr); err != nil {
			return
		}
		newsalt = saltstr
	} else {
		if salt, err = randomBytes(saltLen); err != nil {
			return
		}
		newsalt = hex.EncodeToString(salt)
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

type Encrypted[T any] struct {
	Value     *T
	Encrypted []byte
}

func NewEncrypted[T any](Value *T) Encrypted[T] {
	return Encrypted[T]{Value, nil}
}

func (e *Encrypted[T]) Seal(ciph Cipher) (err error) {
	var plaintext []byte
	if plaintext, err = json.Marshal(e.Value); err != nil {
		return
	}
	e.Encrypted, err = ciph.Encrypt(plaintext)
	e.Value = nil
	return
}
func (e *Encrypted[T]) Open(ciph Cipher) (err error) {
	e.Value = new(T)
	var plaintext []byte
	if plaintext, err = ciph.Decrypt(e.Encrypted); err != nil {
		return
	}
	e.Encrypted = nil
	err = json.Unmarshal(plaintext, e.Value)
	return
}

type Cipher interface {
	Encrypt(plaintext []byte) (buf []byte, err error)
	Decrypt(buf []byte) (plaintext []byte, err error)
}

type aesgcmcipher struct {
	noncesize int
	cipher.AEAD
}

func NewAESGCMCipher(key []byte) (c Cipher, err error) {
	var block cipher.Block
	var aesgcm cipher.AEAD
	if block, err = aes.NewCipher(key); err != nil {
		return
	}
	if aesgcm, err = cipher.NewGCM(block); err != nil {
		return
	}
	c = &aesgcmcipher{noncesize: aesgcm.NonceSize(), AEAD: aesgcm}
	return
}

func (c *aesgcmcipher) Encrypt(plaintext []byte) (buf []byte, err error) {
	if plaintext == nil {
		return nil, nil
	}
	buf = make([]byte, c.noncesize+c.Overhead()+len(plaintext))
	nonce := buf[:c.noncesize]
	if _, err = rand.Read(nonce); err != nil {
		return
	}

	sealed := c.Seal(buf[c.noncesize:c.noncesize], nonce, plaintext, nil)
	buf = buf[:c.noncesize+len(sealed)]
	return
}

func (c *aesgcmcipher) Decrypt(buf []byte) (plaintext []byte, err error) {
	if buf == nil {
		return nil, nil
	}
	if len(buf) < c.noncesize+c.Overhead() {
		return nil, fmt.Errorf("ciphertext too short")
	}
	nonce := buf[:c.noncesize]
	ciphertext := buf[c.noncesize:]
	plaintext, err = c.Open(nil, nonce, ciphertext, nil)
	return
}

type nopcipher struct{}

func NewNopCipher() Cipher {
	return &nopcipher{}
}

func (c *nopcipher) Encrypt(plaintext []byte) (buf []byte, err error) {
	return plaintext, nil
}

func (c *nopcipher) Decrypt(buf []byte) (plaintext []byte, err error) {
	return buf, nil
}
