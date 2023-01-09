package client

import (
	"context"
	"errors"
	"fmt"
)

const authHeader string = "authorization"

func (r *CommandRunner) register(ctx context.Context) error {
	username := getInput("username:", notEmpty)
	password := getInputWithMask("password:", notEmpty)
	paswordConfirm := getInputWithMask("password(confirm):", notEmpty)
	if password != paswordConfirm {
		return errors.New("password mismatch")
	}

	req := buildRegisterRequest(username, password)
	_, err := r.client.AuthClient.Register(ctx, req)
	if err != nil {
		return fmt.Errorf("error register user %s", err.Error())
	}

	cmd := inputSelect("sign in:",
		[]string{Login},
	)
	err = r.run(ctx, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (r *CommandRunner) login(ctx context.Context) error {
	username := getInput("username:", notEmpty)
	password := getInputWithMask("password:", notEmpty)

	req := buildLoginRequest(username, password)
	resp, err := r.client.AuthClient.Login(ctx, req)
	if err != nil {
		return fmt.Errorf("error login user %s", err.Error())
	}

	r.storage.SetToken(resp.AccessToken)
	return nil
}
