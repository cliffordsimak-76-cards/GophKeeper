//go:generate rm -rf ./mock_gen.go
//go:generate mockgen -destination=./mock_gen.go -package=crypto -source=crypto.go
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"golang.org/x/crypto/bcrypt"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/config"
)

// Client represents a crypto manager
type Client interface {
	HashAndSalt(string) (string, error)
	IsCorrectPassword(string, string) bool
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
}

type crypto struct {
	cfg *config.Config
}

// NewClient returns a new client
func NewClient(cfg *config.Config) *crypto {
	return &crypto{cfg}
}

// HashAndSalt hashes a given string
func (c *crypto) HashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// IsCorrectPassword checks hash and password strings
func (c *crypto) IsCorrectPassword(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
}

// Encrypt encrypts and authenticates plaintext
func (c *crypto) Encrypt(str string) (string, error) {
	key := []byte(c.cfg.EncryptPassword)

	aesblock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encrypted := aesgcm.Seal(nonce, nonce, []byte(str), nil)

	return string(encrypted), nil
}

// Decrypt decrypts and authenticates ciphertext
func (c *crypto) Decrypt(str string) (string, error) {
	key := []byte(c.cfg.EncryptPassword)
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return "", err
	}

	ciphertext := []byte(str)
	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("error ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	decrypted, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}
