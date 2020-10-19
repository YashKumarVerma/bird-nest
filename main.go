package main

import (
	"github.com/YashKumarVerma/bird-nest/modules/ui"
	"github.com/c-bata/go-prompt"
)

func completer(document prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(suggestions, document.GetWordBeforeCursor(), true)
}

func main() {
	ui.Heading("Bird Nest")
	ui.Description("a handy repl shell to quickly build nest.js services")
	// fmt.Println("Please select table.")
	// t := prompt.Input("> ", completer)
	// fmt.Println("You selected " + t)
}
