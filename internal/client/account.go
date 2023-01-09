package client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/metadata"
)

// Account represents users account entry
type Account struct {
	ID       string
	Name     string
	Login    string
	Password string
	Metadata string
}

func (r *CommandRunner) createAccount(ctx context.Context) error {
	md := metadata.New(map[string]string{authHeader: r.storage.Token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	account := &Account{}
	account.Name = getInput("name:", notEmpty)
	account.Login = getInput("login:", notEmpty)
	account.Password = getInput("password:", notEmpty)
	account.Metadata = getInput("metadata:", any)

	req := buildCreateAccountRequest(account)
	_, err := r.client.AccountClient.CreateAccount(ctx, req)
	if err != nil {
		log.Printf("error create account: %s", err)
		return fmt.Errorf("error create account")
	}

	return nil
}

func (r *CommandRunner) getAccount(ctx context.Context) error {
	md := metadata.New(map[string]string{authHeader: r.storage.Token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	availableAccounts, err := r.listAvailableAccounts(ctx)
	if err != nil {
		log.Printf("error list available accounts: %s", err)
		return fmt.Errorf("error list available accounts")
	}

	if len(availableAccounts.Names) == 0 {
		getInput("you dont have accounts:", any)
		return nil
	}

	name := inputSelect("choose account", availableAccounts.Names)

	getReq := buildGetAccountRequest(availableAccounts.IdByNameMap[name])
	account, err := r.client.AccountClient.GetAccount(ctx, getReq)
	if err != nil {
		log.Printf("error get account: %s", err)
		return fmt.Errorf("error get account")
	}

	log.Println(account)
	return nil
}

func (r *CommandRunner) updateAccount(ctx context.Context) error {
	md := metadata.New(map[string]string{authHeader: r.storage.Token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	availableAccounts, err := r.listAvailableAccounts(ctx)
	if err != nil {
		log.Printf("error list available accounts: %s", err)
		return fmt.Errorf("error list available accounts")
	}

	if len(availableAccounts.Names) == 0 {
		getInput("you dont have accounts:", any)
		return nil
	}

	accountName := inputSelect("choose account", availableAccounts.Names)

	getReq := buildGetAccountRequest(availableAccounts.IdByNameMap[accountName])
	account, err := r.client.AccountClient.GetAccount(ctx, getReq)
	if err != nil {
		log.Printf("error get account: %s", err)
		return fmt.Errorf("error get account")
	}

	newAccount := pbAccountToAccount(account)
	name := getInput(fmt.Sprintf("name [%s]:", account.Name), any)
	if name != "" {
		newAccount.Name = name
	}
	login := getInput(fmt.Sprintf("login [%s]:", account.Login), any)
	if login != "" {
		newAccount.Login = login
	}
	password := getInput(fmt.Sprintf("password [%s]:", account.Password), any)
	if password != "" {
		newAccount.Password = password
	}
	metadata := getInput(fmt.Sprintf("metadata [%s]:", account.Metadata), any)
	if metadata != "" {
		newAccount.Metadata = metadata
	}

	req := buildUpdateAccountRequest(newAccount)
	_, err = r.client.AccountClient.UpdateAccount(ctx, req)
	if err != nil {
		log.Printf("error update account: %s", err)
		return fmt.Errorf("error update account")
	}

	return nil
}

type availableAccounts struct {
	Names       []string
	IdByNameMap map[string]string
}

func (r *CommandRunner) listAvailableAccounts(ctx context.Context) (*availableAccounts, error) {
	result := &availableAccounts{
		Names:       make([]string, 0),
		IdByNameMap: make(map[string]string),
	}

	req := buildListAvailableAccountsRequest()
	resp, err := r.client.AccountClient.ListAvailableAccounts(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(resp.GetAccounts()) == 0 {
		return result, nil
	}

	for _, c := range resp.GetAccounts() {
		result.IdByNameMap[c.Name] = c.Id
		result.Names = append(result.Names, c.Name)
	}

	return result, nil
}
