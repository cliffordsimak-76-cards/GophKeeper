package authservice

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func Test_Login(t *testing.T) {
	t.Run("validation error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.LoginRequest{}

		_, err := te.service.Login(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("repository error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.LoginRequest{
			Username: "username",
			Password: "password",
		}

		te.userRepoMock.EXPECT().Get(te.ctx, "username").
			Return(nil, errAny)

		_, err := te.service.Login(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("equal passwor error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.LoginRequest{
			Username: "username",
			Password: "password",
		}

		user := &model.User{HashedPassword: "123"}
		te.userRepoMock.EXPECT().Get(te.ctx, "username").
			Return(user, nil)

		te.cryptoClientMock.EXPECT().IsCorrectPassword(user.HashedPassword, req.Password).
			Return(false)

		_, err := te.service.Login(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.NotFound, status.Code(err))
	})

	t.Run("error jwt generate", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.LoginRequest{
			Username: "username",
			Password: "password",
		}

		user := &model.User{HashedPassword: "$2a$10$s33aWwgtuepHxtmHN1qsxuu7HsKFp4sTff4mBoutUXxavEqqBREGe"}
		te.userRepoMock.EXPECT().Get(te.ctx, "username").
			Return(user, nil)

		te.cryptoClientMock.EXPECT().IsCorrectPassword(user.HashedPassword, req.Password).
			Return(true)

		te.jwtClientMock.EXPECT().Generate(user).
			Return("", errAny)

		_, err := te.service.Login(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.LoginRequest{
			Username: "username",
			Password: "password",
		}

		user := &model.User{HashedPassword: "$2a$10$s33aWwgtuepHxtmHN1qsxuu7HsKFp4sTff4mBoutUXxavEqqBREGe"}
		te.userRepoMock.EXPECT().Get(te.ctx, "username").
			Return(user, nil)

		te.cryptoClientMock.EXPECT().IsCorrectPassword(user.HashedPassword, req.Password).
			Return(true)

		token := "token"
		te.jwtClientMock.EXPECT().Generate(user).
			Return(token, nil)

		response, err := te.service.Login(te.ctx, req)
		require.NoError(t, err)
		require.Equal(t, &api.LoginResponse{AccessToken: token}, response)
	})
}
