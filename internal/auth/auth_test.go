package auth

import (
	"context"
	"testing"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/config"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func Test_GetUserIdFromContext(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authImpl := NewAuth(&config.Config{SecretKey: "secret"})

		testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHA" +
			"iOjE2NzIyNDkzODIsInVzZXJJRCI6IjhhMmVjMzk3LTIyZDAtNGQyMS1hOTE3L" +
			"WJhODI4MGQ5N2NlNiJ9.IsU6Z9SLWy7NwpNCCioTBRMCVc3FBq4zylr-yAglitY"
		ctx := context.Background()
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("authorization", testToken),
		)

		expectedUserID := "8a2ec397-22d0-4d21-a917-ba8280d97ce6"
		userID, err := authImpl.GetUserIdFromContext(ctx)
		require.NoError(t, err)
		require.Equal(t, expectedUserID, userID)
	})
}

func Test_getTokenFromContext(t *testing.T) {
	t.Run("error metadata is not provided", func(t *testing.T) {
		ctx := context.Background()

		_, err := getTokenFromContext(ctx)

		require.Error(t, err)
		require.Equal(t, codes.Unauthenticated, status.Code(err))
	})
	t.Run("error authorization token is not provided", func(t *testing.T) {
		ctx := context.Background()

		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("some key", "some value"),
		)

		_, err := getTokenFromContext(ctx)

		require.Error(t, err)
		require.Equal(t, codes.Unauthenticated, status.Code(err))
	})
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHA" +
			"iOjE2NzIyNDkzODIsInVzZXJJRCI6IjhhMmVjMzk3LTIyZDAtNGQyMS1hOTE3L" +
			"WJhODI4MGQ5N2NlNiJ9.IsU6Z9SLWy7NwpNCCioTBRMCVc3FBq4zylr-yAglitY"
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("authorization", testToken),
		)

		token, err := getTokenFromContext(ctx)
		require.NoError(t, err)
		require.Equal(t, testToken, token)
	})
}
