package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/accountservice"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/authservice"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/cardservice"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/noteservice"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/auth"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/config"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/crypto"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/db"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/jwt"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

const (
	clientCACertFile = "./cert/ca-cert.pem"
	serverCertFile   = "./cert/server-cert.pem"
	serverKeyFile    = "./cert/server-key.pem"
)

func Run(ctx context.Context, cfg *config.Config) error {
	env := initEnv(ctx, cfg)

	db.Ping(ctx, env.db, env.cfg)

	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(env.auth.Unary()),
	}

	if cfg.EnableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			return fmt.Errorf("cannot load TLS credentials: %w", err)
		}

		serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))
	}

	s := grpc.NewServer(serverOptions...)

	authService := authservice.NewService(env.repoGroup, env.jwt, env.crypto)
	cardService := cardservice.NewService(env.repoGroup, env.auth, env.crypto)
	accountService := accountservice.NewService(env.repoGroup, env.auth)
	noteService := noteservice.NewService(env.repoGroup, env.auth)

	api.RegisterAuthServiceServer(s, authService)
	api.RegisterAccountServiceServer(s, accountService)
	api.RegisterCardServiceServer(s, cardService)
	api.RegisterNoteServiceServer(s, noteService)

	listener, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("failed to listen tcp %s, %v", cfg.GRPCAddr, err)
	}

	log.Printf("starting listening grpc server at %s", cfg.GRPCAddr)
	return s.Serve(listener)
}

type env struct {
	db        *sqlx.DB
	cfg       *config.Config
	repoGroup *repository.Group
	jwt       jwt.Client
	auth      auth.Client
	crypto    crypto.Client
}

func initEnv(ctx context.Context, cfg *config.Config) *env {
	db, err := db.NewClient(ctx, cfg)
	if err != nil {
		log.Fatal("error connect to db ", err)
	}

	repoGroup := repository.NewGroup(db)

	jwt := jwt.NewClient(cfg)

	auth := auth.NewClient(jwt)

	crypto := crypto.NewClient(cfg)

	return &env{
		db:        db,
		cfg:       cfg,
		repoGroup: repoGroup,
		jwt:       jwt,
		auth:      auth,
		crypto:    crypto,
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed client's certificate
	pemClientCA, err := os.ReadFile(clientCACertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}
