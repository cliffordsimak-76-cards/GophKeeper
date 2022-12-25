package client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/metadata"
)

// Card represents users Credit card entry
type Card struct {
	Name   string
	Number string
	Expire string
	CVC    string
	Holder string
}

func (r *CommandRunner) createCard(ctx context.Context) error {
	md := metadata.New(map[string]string{authHeader: r.storage.Token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	card := &Card{}
	card.Name = getInput("name:", notEmpty)
	card.Number = getInput("number:", notEmpty)
	card.Expire = getInput("expire:", notEmpty)
	card.CVC = getInput("CVC:", notEmpty)
	card.Holder = getInput("holder:", notEmpty)

	req := buildCreateCardRequest(card)
	_, err := r.client.CardClient.Create(ctx, req)
	if err != nil {
		log.Printf("error create card: %s", err)
		return fmt.Errorf("error create card")
	}

	return nil
}

func (r *CommandRunner) getCard(ctx context.Context) error {
	md := metadata.New(map[string]string{authHeader: r.storage.Token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	availableCards, err := r.listAvailableCards(ctx)
	if err != nil {
		log.Printf("error list available cards: %s", err)
		return fmt.Errorf("error list available cards")
	}

	if len(availableCards.Names) == 0 {
		getInput("you dont have cards:", any)
		return nil
	}

	name := inputSelect("choose card", availableCards.Names)

	getReq := buildGetCardRequest(availableCards.IdByNameMap[name])
	card, err := r.client.CardClient.Get(ctx, getReq)
	if err != nil {
		log.Printf("error get card: %s", err)
		return fmt.Errorf("error get card")
	}

	log.Println(card)
	return nil
}

func (r *CommandRunner) updateCard(ctx context.Context) error {
	md := metadata.New(map[string]string{authHeader: r.storage.Token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	availableCards, err := r.listAvailableCards(ctx)
	if err != nil {
		log.Printf("error list available cards: %s", err)
		return fmt.Errorf("error list available cards")
	}

	if len(availableCards.Names) == 0 {
		getInput("you dont have cards:", any)
		return nil
	}

	cardName := inputSelect("choose card", availableCards.Names)

	getReq := buildGetCardRequest(availableCards.IdByNameMap[cardName])
	card, err := r.client.CardClient.Get(ctx, getReq)
	if err != nil {
		log.Printf("error get card: %s", err)
		return fmt.Errorf("error get card")
	}

	newCard := pbCardToCard(card)
	name := getInput(fmt.Sprintf("name [%s]:", card.Name), notEmpty)
	if name != "" {
		newCard.Name = name
	}
	number := getInput(fmt.Sprintf("number [%s]:", card.Number), notEmpty)
	if number != "" {
		newCard.Number = number
	}
	expire := getInput(fmt.Sprintf("expire [%s]:", card.Expire), notEmpty)
	if expire != "" {
		newCard.Expire = expire
	}
	cvc := getInput(fmt.Sprintf("cvc [%s]:", card.Cvc), notEmpty)
	if cvc != "" {
		newCard.CVC = cvc
	}
	holder := getInput(fmt.Sprintf("holder [%s]:", card.Holder), notEmpty)
	if holder != "" {
		newCard.Holder = holder
	}

	req := buildUpdateCardRequest(newCard)
	_, err = r.client.CardClient.Update(ctx, req)
	if err != nil {
		log.Printf("error update card: %s", err)
		return fmt.Errorf("error update card")
	}

	return nil
}

type availableCards struct {
	Names       []string
	IdByNameMap map[string]string
}

func (r *CommandRunner) listAvailableCards(ctx context.Context) (*availableCards, error) {
	result := &availableCards{
		Names:       make([]string, 0),
		IdByNameMap: make(map[string]string),
	}

	req := buildListAvailableCardsRequest()
	resp, err := r.client.CardClient.ListAvailableCards(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(resp.GetCards()) == 0 {
		return result, nil
	}

	for _, c := range resp.GetCards() {
		result.IdByNameMap[c.Name] = c.Id
		result.Names = append(result.Names, c.Name)
	}

	return result, nil
}
