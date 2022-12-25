package client

import "context"

const (
	Register string = "Register"
	Login    string = "Login"

	CreateCard string = "Create card"
	GetCard    string = "Get card"
	UpdateCard string = "Update card"
)

type CommandRunner struct {
	client  *Client
	storage *MemStorage
}

func NewCommandRunner(
	client *Client,
	storage *MemStorage,
) *CommandRunner {
	return &CommandRunner{
		client:  client,
		storage: storage,
	}
}

func (r *CommandRunner) run(
	ctx context.Context,
	cmd string,
) error {
	switch cmd {
	case Register:
		return r.register(ctx)
	case Login:
		return r.login(ctx)
	case CreateCard:
		return r.createCard(ctx)
	case GetCard:
		return r.getCard(ctx)
	case UpdateCard:
		return r.updateCard(ctx)
	}

	return nil
}
