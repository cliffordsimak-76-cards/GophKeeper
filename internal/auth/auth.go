//go:generate rm -rf ./mock_gen.go
//go:generate mockgen -destination=./mock_gen.go -package=auth -source=auth.go
package auth

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/jwt"
)

var (
	userServiceRegister = "/gophkeeper.v1.AuthService/Register"
	userServiceLogin    = "/gophkeeper.v1.AuthService/Login"
	ignoreMethod        = []string{userServiceRegister, userServiceLogin}
	authHeader          = "authorization"
)

// Client represents a auth manager
type Client interface {
	ExtractUserIdFromContext(ctx context.Context) (string, error)
	Unary() grpc.UnaryServerInterceptor
}

// client is a server interceptor for authentication and authorization
type client struct {
	jwt jwt.Client
}

// NewClient returns a new auth interceptor
func NewClient(jwt jwt.Client) *client {
	return &client{jwt}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func (c *client) Unary() grpc.UnaryServerInterceptor {
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

		err := c.authorize(ctx)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (c *client) authorize(ctx context.Context) error {
	accessToken, err := extractTokenFromContext(ctx)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "error authorize: %v", err)
	}
	err = c.jwt.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return nil
}

// ExtractUserIdFromContext extracts userID from authorization context header
func (c *client) ExtractUserIdFromContext(ctx context.Context) (string, error) {
	accessToken, err := extractTokenFromContext(ctx)
	if err != nil {
		return "", fmt.Errorf("error get token from context: %w", err)
	}

	userID, err := c.jwt.ExtractUserID(accessToken)
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
