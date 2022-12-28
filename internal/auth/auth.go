package auth

import (
	"context"
	"fmt"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/config"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthImpl struct {
	secretKey string
}

// NewAuth returns a new auth manager
func NewAuth(cfg *config.Config) *AuthImpl {
	return &AuthImpl{cfg.SecretKey}
}

func (a *AuthImpl) GetUserIdFromContext(ctx context.Context) (string, error) {
	tokenString, err := getTokenFromContext(ctx)
	if err != nil {
		return "", fmt.Errorf("error get token from context: %w", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}

		return []byte(a.secretKey), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["userID"].(string), nil
	} else {
		return "", fmt.Errorf("no userID in token")
	}
}

func getTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	return values[0], nil
}
