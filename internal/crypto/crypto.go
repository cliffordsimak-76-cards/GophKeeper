//go:generate rm -rf ./mock_gen.go
//go:generate mockgen -destination=./mock_gen.go -package=crypto -source=crypto.go
package crypto

import "golang.org/x/crypto/bcrypt"

type Crypto interface {
	HashAndSalt(pwd string) (string, error)
	IsCorrectPassword(hashedPwd string, plainPwd string) bool
}

type CryptoImpl struct{}

// HashAndSalt Hashes a given string
func (c *CryptoImpl) HashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (c *CryptoImpl) IsCorrectPassword(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
}
