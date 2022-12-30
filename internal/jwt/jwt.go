//go:generate rm -rf ./mock_gen.go
//go:generate mockgen -destination=./mock_gen.go -package=jwt -source=jwt.go
package jwt

import (
	"fmt"
	"time"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/config"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	"github.com/golang-jwt/jwt"
)

// JWT represents a token manager
type JWT interface {
	Generate(user *model.User) (string, error)
	Verify(accessToken string) error
	ExtractUserID(accessToken string) (string, error)
}

// JWTImpl is a web token manager
type JWTImpl struct {
	secretKey     string
	tokenDuration time.Duration
}

// NewJWTImpl returns a new JWTImpl
func NewJWTImpl(cfg *config.Config) *JWTImpl {
	return &JWTImpl{cfg.SecretKey, cfg.TokenDuration}
}

// UserClaims is a custom JWTImpl claims that contains some user's information
type userClaims struct {
	jwt.StandardClaims
	UserID string `json:"userID"`
}

// Generate generates and signs a new token for a user
func (j *JWTImpl) Generate(user *model.User) (string, error) {
	claims := userClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.tokenDuration).Unix(),
		},
		UserID: user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// Verify verifies the access token string and return a user claim if the token is valid
func (j *JWTImpl) Verify(accessToken string) error {
	_, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})
	if err != nil {
		return fmt.Errorf("error parse jwt token: %w", err)
	}

	return nil
}

// ExtractUserID extracts user identifier from access token string
func (j *JWTImpl) ExtractUserID(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})
	if err != nil {
		return "", fmt.Errorf("error parse jwt token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("error parse token claims")
	}

	if userID, ok := claims["userID"].(string); ok {
		return userID, nil
	}

	return "", fmt.Errorf("error no userID in token")
}
