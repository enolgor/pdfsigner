package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/goccy/go-json"
	"github.com/goforj/godump"
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

type Encrypted[T any] struct {
	Value     *T
	Encrypted []byte
}

func NewEncrypted[T any](Value *T) Encrypted[T] {
	return Encrypted[T]{Value, nil}
}

func (e *Encrypted[T]) Seal(ciph *Cipher) (err error) {
	var plaintext []byte
	if plaintext, err = json.Marshal(e.Value); err != nil {
		return
	}
	e.Encrypted, err = ciph.Encrypt(plaintext)
	e.Value = nil
	return
}
func (e *Encrypted[T]) Open(ciph *Cipher) (err error) {
	e.Value = new(T)
	var plaintext []byte
	if plaintext, err = ciph.Decrypt(e.Encrypted); err != nil {
		return
	}
	e.Encrypted = nil
	err = json.Unmarshal(plaintext, e.Value)
	return
}

type MyStruct struct {
	Name string `json:"name"`
}

func main() {
	password := "3cerdit0s"
	key, _, err := Derive(password, "")
	if err != nil {
		panic(err)
	}
	data := MyStruct{
		Name: "Eneko",
	}
	ciph, err := NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted := NewEncrypted(&data)
	if err = encrypted.Seal(ciph); err != nil {
		panic(err)
	}
	encoded, err := json.Marshal(encrypted)
	if err != nil {
		panic(err)
	}
	godump.Dump(encoded)
	if err := json.Unmarshal(encoded, &encrypted); err != nil {
		panic(err)
	}
	if err = encrypted.Open(ciph); err != nil {
		panic(err)
	}
	godump.Dump(encrypted)
}

type Cipher struct {
	noncesize int
	cipher.AEAD
}

func NewCipher(key []byte) (c *Cipher, err error) {
	var block cipher.Block
	var aesgcm cipher.AEAD
	if block, err = aes.NewCipher(key); err != nil {
		return
	}
	if aesgcm, err = cipher.NewGCM(block); err != nil {
		return
	}
	c = &Cipher{noncesize: aesgcm.NonceSize(), AEAD: aesgcm}
	return
}

func (c *Cipher) Encrypt(plaintext []byte) (buf []byte, err error) {
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

func (c *Cipher) Decrypt(buf []byte) (plaintext []byte, err error) {
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
