package shell

import (
	"fmt"
	"strings"

	"github.com/YashKumarVerma/bird-nest/modules/entity"

	"github.com/c-bata/go-prompt"
	"github.com/kyokomi/emoji"
)

// Init the shell to take input from user
func Init() {
	welcomeMessage := emoji.Sprintf("Welcome to the :tiger: shell.\n This shell has autocomplete to make your work easy !")
	fmt.Println(welcomeMessage)

	columns := make([]string, 0)
	for true {
		command := prompt.Input(" > ", entity.AutoComplete,
			prompt.OptionTitle("Bird Nest"),
			prompt.OptionHistory([]string{"exit"}),
			prompt.OptionPrefixTextColor(prompt.Yellow),
			prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
			prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
			prompt.OptionSuggestionBGColor(prompt.DarkGray))
		if command == "exit" {
			break
		} else {
			if strings.TrimSpace(command) != "" {
				columns = append(columns, strings.TrimSpace(command))
			}
		}
	}

	entity.Process(columns)

}
