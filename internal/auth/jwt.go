package auth

import (
	"fmt"
	"time"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/config"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	"github.com/golang-jwt/jwt"
)

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
	_, err := jwt.ParseWithClaims(
		accessToken,
		&userClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(j.secretKey), nil
		},
	)
	if err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}

	return nil
}

// func (j *JWTImpl) GetUserIdFromToken(tokenString string) (string, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return []byte(j.secretKey), nil
// 	})

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		return claims["userID"].(string), nil
// 	} else {
// 		return "", err
// 	}
// }
