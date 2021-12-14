package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/theovassiliou/cipher"
)

type promptContent struct {
	errorMsg string
	label    string
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func promptGetSelect(pc promptContent) string {
	items := []string{"Keyword Cipher", "Rotate Cipher"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

type ciph struct {
	Name        string
	Description string
	Codec       cipher.CiphererDecipherer
}

func main() {

	kws := cipher.NewKeywordCipherer("", cipher.StdLowercaseAlphabet)
	cs := cipher.NewCaesarCipher(cipher.StdLowercaseAlphabet)

	ciphers := []ciph{
		{
			Name:        kws.Name(),
			Description: kws.Description(),
			Codec:       kws,
		},
		{
			Name:        cs.Name(),
			Description: cs.Description(),
			Codec:       cs,
		},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "> {{ .Name | cyan }} ({{ .Name | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .Name | red }})",
		Selected: "> {{ .Name | red | cyan }}",
		Details: `
----------- Ciphers -----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Description:" | faint }}	
{{ .Description }}`,
	}

	searcher := func(input string, index int) bool {
		c := ciphers[index]
		name := strings.Replace(strings.ToLower(c.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Ciphers",
		Items:     ciphers,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	textPromptContent := promptContent{
		"Please enter the plain text.",
		"Plain text:",
	}

	text := promptGetInput(textPromptContent)

	fmt.Println(ciphers[i].Codec.Cipher(text))
}
