package client

import (
	"context"
	"log"
)

func Run(ctx context.Context, cfg *Config) error {
	client, err := NewClient(cfg)
	if err != nil {
		log.Fatalf("error create client %v", err)
	}

	storage := &MemStorage{}

	runner := NewCommandRunner(client, storage)

	cmd := inputSelect("sing up or sign in:",
		[]string{Register, Login},
	)
	err = runner.run(ctx, cmd)
	if err != nil {
		return err
	}

	for {
		cmd := inputSelect("choose action:",
			[]string{
				CreateCard,
				GetCard,
			},
		)
		err = runner.run(ctx, cmd)
		if err != nil {
			return err
		}
	}
}
