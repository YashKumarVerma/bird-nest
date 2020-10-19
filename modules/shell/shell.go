package shell

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/kyokomi/emoji"
)

func completer(document prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "create", Description: "create a new resource"},
		{Text: "entity", Description: "create a new entity and describe database storage"},
	}
	return prompt.FilterHasPrefix(suggestions, document.GetWordBeforeCursor(), true)
}

// Init the shell to take input from user
func Init() {
	welcomeMessage := emoji.Sprint("Welcome to the :tiger: shell")
	fmt.Println(welcomeMessage)
	command := prompt.Input("> ", completer)
	fmt.Println("You selected " + command)
}
