package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/YashKumarVerma/bird-nest/modules/entity"
	"github.com/YashKumarVerma/bird-nest/modules/generators"
	"github.com/c-bata/go-prompt"
	"github.com/kyokomi/emoji"
)

// Init the shell to take input from user
func Init() {
	fmt.Print("Enter module name : ")
	reader := bufio.NewReader(os.Stdin)
	moduleName, _ := reader.ReadString('\n')
	moduleName = strings.Trim(moduleName, " ")
	moduleName = strings.Trim(moduleName, "\n")
	// moduleName := "user"

	welcomeMessage := emoji.Sprintf("\n\nSimply press space key to show suggestions")
	fmt.Println(welcomeMessage)

	var entitySchemas []entity.StructuredCommandData
	columns := make([]string, 0)
	commandHistory := make([]string, 0)
	commandHistory = append(commandHistory, "exit")
	for true {
		command := prompt.Input(" > ", entity.AutoComplete,
			prompt.OptionTitle("Bird Nest"),
			prompt.OptionHistory(commandHistory),
			prompt.OptionPrefixTextColor(prompt.Yellow),
			prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
			prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
			prompt.OptionSuggestionBGColor(prompt.DarkGray))
		if command == "exit" {
			break
		} else {
			if strings.TrimSpace(command) != "" {
				commandHistory = append(commandHistory, strings.TrimSpace(command))
				columns = append(columns, strings.TrimSpace(command))
			}
		}
	}

	// columns = append(columns, "--name:id --type:int --auto_increment --primary")
	// columns = append(columns, "--name:name --type:string --unique")
	// columns = append(columns, "--name:age --type:int")
	// columns = append(columns, "--name:money --type:int --default:0")
	// columns = append(columns, "--name:gender --type:string")
	// columns = append(columns, "--name:devices --type:string --array")
	// columns = append(columns, "--name:country --type:string --default:in --length:2")
	// columns = append(columns, "--name:github --type:string --unique")
	// columns = append(columns, "--name:email --type:string --unique")
	// columns = append(columns, "--name:instagram --type:string --unique")

	entitySchemas = entity.Process(columns)
	generators.Init(moduleName, entitySchemas)
}
