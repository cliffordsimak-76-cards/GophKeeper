package authservice

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func Test_Register(t *testing.T) {
	t.Run("validation error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.RegisterRequest{}

		_, err := te.service.Register(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("crypto error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.RegisterRequest{
			Username: "username",
			Password: "password",
		}

		te.cryptoClientMock.EXPECT().HashAndSalt("password").
			Return("", errAny)

		_, err := te.service.Register(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("repository error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.RegisterRequest{
			Username: "username",
			Password: "password",
		}

		hashedPwd := "hashedPwd"
		te.cryptoClientMock.EXPECT().HashAndSalt("password").
			Return(hashedPwd, nil)

		user := buildUser("username", hashedPwd)
		te.userRepoMock.EXPECT().Create(te.ctx, user).
			Return(nil, errAny)

		_, err := te.service.Register(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.RegisterRequest{
			Username: "username",
			Password: "password",
		}

		hashedPwd := "hashedPwd"
		te.cryptoClientMock.EXPECT().HashAndSalt(req.Password).
			Return(hashedPwd, nil)

		user := buildUser("username", hashedPwd)
		te.userRepoMock.EXPECT().Create(te.ctx, user).
			Return(&model.User{}, nil)

		response, err := te.service.Register(te.ctx, req)
		require.NoError(t, err)
		require.Equal(t, &api.RegisterResponse{}, response)
	})
}

func Test_buildUser(t *testing.T) {
	t.Run("all data filler", func(t *testing.T) {
		username := "username"
		password := "password"
		user := buildUser(username, password)
		expectedUser := &model.User{
			Username:       username,
			HashedPassword: password,
		}
		require.Equal(t, expectedUser, user)
	})
}
