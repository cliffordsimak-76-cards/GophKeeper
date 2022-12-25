package auth

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	userServiceRegister = "/gophkeeper.v1.AuthService/Register"
	userServiceLogin    = "/gophkeeper.v1.AuthService/Login"
	ignoreMethod        = []string{userServiceRegister, userServiceLogin}
	// userIDClaim         = "userID"
)

// AuthInterceptor is a server interceptor for authentication and authorization
type AuthInterceptor struct {
	jwt JWT
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(jwt JWT) *AuthInterceptor {
	return &AuthInterceptor{jwt}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
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

		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

// Stream returns a server interceptor function to authenticate and authorize stream RPC
func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("--> stream interceptor: ", info.FullMethod)

		err := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, stream)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	accessToken, err := getTokenFromContext(ctx)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "error authorize: %v", err)
	}
	err = interceptor.jwt.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return nil
}
