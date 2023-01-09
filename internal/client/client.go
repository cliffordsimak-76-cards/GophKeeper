//go:generate rm -rf ./mock_gen.go
//go:generate mockgen -destination=./mock_gen.go -package=client -source=client.go
package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverCACertFile = "./cert/ca-cert.pem"
	clientCertFile   = "./cert/client-cert.pem"
	clientKeyFile    = "./cert/client-key.pem"
)

type Services interface {
	api.AuthServiceClient
	api.CardServiceClient
	api.AccountServiceClient
	api.NoteServiceClient
}

type Client struct {
	AuthClient    api.AuthServiceClient
	CardClient    api.CardServiceClient
	AccountClient api.AccountServiceClient
	NoteClient    api.NoteServiceClient
}

func NewClient(cfg *Config) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ServerTimeout)
	defer cancel()

	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())

	if cfg.EnableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}

		transportOption = grpc.WithTransportCredentials(tlsCredentials)
	}

	conn, err := grpc.DialContext(
		ctx,
		cfg.ServerHost,
		transportOption,
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		AuthClient:    api.NewAuthServiceClient(conn),
		CardClient:    api.NewCardServiceClient(conn),
		AccountClient: api.NewAccountServiceClient(conn),
		NoteClient:    api.NewNoteServiceClient(conn),
	}, nil
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := os.ReadFile(serverCACertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}
