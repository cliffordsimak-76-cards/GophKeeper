package client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/metadata"
)

// Note represents users note entry
type Note struct {
	ID       string
	Name     string
	Text     string
	Metadata string
}

func (r *CommandRunner) createNote(ctx context.Context) error {
	md := metadata.New(map[string]string{authHeader: r.storage.Token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	note := &Note{}
	note.Name = getInput("name:", notEmpty)
	note.Text = getInput("text:", notEmpty)
	note.Metadata = getInput("metadata:", any)

	req := buildCreateNoteRequest(note)
	_, err := r.client.NoteClient.CreateNote(ctx, req)
	if err != nil {
		log.Printf("error create note: %s", err)
		return fmt.Errorf("error create note")
	}

	return nil
}

func (r *CommandRunner) getNote(ctx context.Context) error {
	md := metadata.New(map[string]string{authHeader: r.storage.Token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	availableNotes, err := r.listAvailableNotes(ctx)
	if err != nil {
		log.Printf("error list available notes: %s", err)
		return fmt.Errorf("error list available notes")
	}

	if len(availableNotes.Names) == 0 {
		getInput("you dont have notes:", any)
		return nil
	}

	name := inputSelect("choose note", availableNotes.Names)

	getReq := buildGetNoteRequest(availableNotes.IdByNameMap[name])
	note, err := r.client.NoteClient.GetNote(ctx, getReq)
	if err != nil {
		log.Printf("error get note: %s", err)
		return fmt.Errorf("error get note")
	}

	log.Println(note)
	return nil
}

func (r *CommandRunner) updateNote(ctx context.Context) error {
	md := metadata.New(map[string]string{authHeader: r.storage.Token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	availableNotes, err := r.listAvailableNotes(ctx)
	if err != nil {
		log.Printf("error list available notes: %s", err)
		return fmt.Errorf("error list available notes")
	}

	if len(availableNotes.Names) == 0 {
		getInput("you dont have notes:", any)
		return nil
	}

	noteName := inputSelect("choose note", availableNotes.Names)

	getReq := buildGetNoteRequest(availableNotes.IdByNameMap[noteName])
	note, err := r.client.NoteClient.GetNote(ctx, getReq)
	if err != nil {
		log.Printf("error get note: %s", err)
		return fmt.Errorf("error get note")
	}

	newNote := pbNoteToNote(note)
	name := getInput(fmt.Sprintf("name [%s]:", note.Name), any)
	if name != "" {
		newNote.Name = name
	}
	text := getInput(fmt.Sprintf("text [%s]:", note.Text), any)
	if text != "" {
		newNote.Text = text
	}
	metadata := getInput(fmt.Sprintf("metadata [%s]:", note.Metadata), any)
	if metadata != "" {
		newNote.Metadata = metadata
	}

	req := buildUpdateNoteRequest(newNote)
	_, err = r.client.NoteClient.UpdateNote(ctx, req)
	if err != nil {
		log.Printf("error update note: %s", err)
		return fmt.Errorf("error update note")
	}

	return nil
}

type availableNotes struct {
	Names       []string
	IdByNameMap map[string]string
}

func (r *CommandRunner) listAvailableNotes(ctx context.Context) (*availableNotes, error) {
	result := &availableNotes{
		Names:       make([]string, 0),
		IdByNameMap: make(map[string]string),
	}

	req := buildListAvailableNotesRequest()
	resp, err := r.client.NoteClient.ListAvailableNotes(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(resp.GetNotes()) == 0 {
		return result, nil
	}

	for _, c := range resp.GetNotes() {
		result.IdByNameMap[c.Name] = c.Id
		result.Names = append(result.Names, c.Name)
	}

	return result, nil
}
