//go:generate rm -rf ./mock_gen.go
//go:generate mockgen -destination=./mock_gen.go -package=auth -source=auth.go
package auth

import (
	"context"
	"fmt"
	"log"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	userServiceRegister = "/gophkeeper.v1.AuthService/Register"
	userServiceLogin    = "/gophkeeper.v1.AuthService/Login"
	ignoreMethod        = []string{userServiceRegister, userServiceLogin}
	authHeader          = "authorization"
)

type Auth interface {
	ExtractUserIdFromContext(ctx context.Context) (string, error)
}

// AuthImpl is a server interceptor for authentication and authorization
type AuthImpl struct {
	jwt jwt.JWT
}

// NewAuthImpl returns a new auth interceptor
func NewAuthImpl(jwt jwt.JWT) *AuthImpl {
	return &AuthImpl{jwt}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func (i *AuthImpl) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		method, _ := grpc.Method(ctx)
		for _, imethod := range ignoreMethod {
			if method == imethod {
				return handler(ctx, req)
			}
		}

		err := i.authorize(ctx)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (i *AuthImpl) authorize(ctx context.Context) error {
	accessToken, err := extractTokenFromContext(ctx)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "error authorize: %v", err)
	}
	err = i.jwt.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return nil
}

// GetUserIdFromContext extracts userID from authorization context header
func (i *AuthImpl) ExtractUserIdFromContext(ctx context.Context) (string, error) {
	accessToken, err := extractTokenFromContext(ctx)
	if err != nil {
		return "", fmt.Errorf("error get token from context: %w", err)
	}

	userID, err := i.jwt.ExtractUserID(accessToken)
	if err != nil {
		return "", fmt.Errorf("error extract userID from token: %w", err)
	}

	return userID, nil
}

func extractTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md[authHeader]
	if len(values) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	return values[0], nil
}
