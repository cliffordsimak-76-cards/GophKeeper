package app

import (
	"context"
	"log"
	"net"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/accountservice"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/authservice"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/cardservice"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/auth"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/config"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/crypto"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/db"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, cfg *config.Config) error {
	env := initEnv(ctx, cfg)

	interceptor := auth.NewAuthInterceptor(env.jwt)
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}
	s := grpc.NewServer(serverOptions...)

	authService := authservice.NewService(env.repoGroup, env.jwt, env.crypto)
	accountService := accountservice.NewService(env.repoGroup, env.auth)
	cardService := cardservice.NewService(env.repoGroup, env.auth)

	api.RegisterAuthServiceServer(s, authService)
	api.RegisterAccountServiceServer(s, accountService)
	api.RegisterCardServiceServer(s, cardService)

	listener, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("failed to listen tcp %s, %v", cfg.GRPCAddr, err)
	}

	log.Printf("starting listening grpc server at %s", cfg.GRPCAddr)
	return s.Serve(listener)
}

type Env struct {
	db        *db.ClientImpl
	cfg       *config.Config
	repoGroup *repository.Group
	jwt       *auth.JWTImpl
	auth      *auth.AuthImpl
	crypto    *crypto.CryptoImpl
}

func initEnv(ctx context.Context, cfg *config.Config) *Env {
	dbClient, err := db.NewClient(cfg)
	if err != nil {
		log.Fatal("not connect to db ", err)
	}

	repoGroup := repository.NewGroup(dbClient)

	jwt := auth.NewJWTImpl(cfg)

	auth := auth.NewAuth(cfg)

	crypto := &crypto.CryptoImpl{}

	return &Env{
		db:        dbClient,
		cfg:       cfg,
		repoGroup: repoGroup,
		jwt:       jwt,
		auth:      auth,
		crypto:    crypto,
	}
}
