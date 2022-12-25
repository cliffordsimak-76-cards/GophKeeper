package authservice

import (
	"context"
	"log"

	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Login(
	ctx context.Context,
	req *api.LoginRequest,
) (*api.LoginResponse, error) {
	err := validateLoginRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.repoGroup.UserRepository.Get(ctx, req.GetUsername())
	if err != nil {
		log.Printf("error get user in db: %s", err)
		return nil, status.Error(codes.Internal, "error get user in db")
	}

	if !s.crypto.IsCorrectPassword(user.HashedPassword, req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := s.jwt.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	return &api.LoginResponse{AccessToken: token}, nil
}

func validateLoginRequest(req *api.LoginRequest) error {
	return validation.Errors{
		"username": validation.Validate(req.GetUsername(), validation.Required),
		"password": validation.Validate(req.GetPassword(), validation.Required),
	}.Filter()
}
