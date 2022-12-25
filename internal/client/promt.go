package client

import (
	"log"

	"github.com/manifoldco/promptui"
)

var templates = &promptui.PromptTemplates{
	Prompt:  "{{ . }} ",
	Valid:   "{{ . | green }} ",
	Invalid: "{{ . | red }} ",
	Success: "{{ . | bold }} ",
}

func inputSelect(label string, items []string) string {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	return selectRun(prompt)
}

func getInput(label string, validator validator) string {
	prompt := promptui.Prompt{
		Label:     label,
		Templates: templates,
		Validate:  promptui.ValidateFunc(validator),
	}

	return promtRun(prompt)
}

func getInputWithMask(label string, validator validator) string {
	prompt := promptui.Prompt{
		Label:     label,
		Templates: templates,
		Validate:  promptui.ValidateFunc(validator),
		Mask:      '*',
		
	}

	return promtRun(prompt)
}

func promtRun(prompt promptui.Prompt) string {
	result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result
}

func selectRun(prompt promptui.Select) string {
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result
}
