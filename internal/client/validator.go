package client

import "errors"

type validator func(string) error

func notEmpty(input string) error {
	if len(input) == 0 {
		return errors.New("")
	}
	return nil
}

func any(input string) error {
	return nil
}