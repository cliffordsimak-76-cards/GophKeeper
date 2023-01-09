package authservice

import (
	"context"
	"log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func (s *Service) Register(
	ctx context.Context,
	req *api.RegisterRequest,
) (*api.RegisterResponse, error) {
	err := validateRegisterRequest(req)
	if err != nil {
		log.Printf("error validate register request: %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	hashedPwd, err := s.crypto.HashAndSalt(req.GetPassword())
	if err != nil {
		log.Printf("error hash password: %s", err)
		return nil, status.Error(codes.Internal, "error hash password")
	}

	user := buildUser(req.GetUsername(), hashedPwd)
	_, err = s.repoGroup.UserRepository.Create(ctx, user)
	if err != nil {
		log.Printf("error create user in db: %s", err)
		return nil, status.Error(codes.Internal, "error create user in db")
	}

	return &api.RegisterResponse{}, nil
}

func validateRegisterRequest(req *api.RegisterRequest) error {
	return validation.Errors{
		"username": validation.Validate(req.GetUsername(), validation.Required),
		"password": validation.Validate(req.GetPassword(), validation.Required),
	}.Filter()
}

func buildUser(username, password string) *model.User {
	return &model.User{
		Username:       username,
		HashedPassword: password,
	}
}
