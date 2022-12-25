//go:generate rm -rf ./mock_gen.go
//go:generate mockgen -destination=./mock_gen.go -package=client -source=client.go
package client

import (
	"context"

	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Services interface {
	api.AuthServiceClient
	api.CardServiceClient
}

type Client struct {
	AuthClient api.AuthServiceClient
	CardClient api.CardServiceClient
}

func NewClient(cfg *Config) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ServerTimeout)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		cfg.ServerHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		AuthClient: api.NewAuthServiceClient(conn),
		CardClient: api.NewCardServiceClient(conn),
	}, nil
}
