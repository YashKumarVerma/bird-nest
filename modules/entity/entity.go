package entity

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/YashKumarVerma/bird-nest/modules/ui"
	"github.com/c-bata/go-prompt"
)

type structuredCommandData struct {
	name          string
	primary       bool
	datatype      string
	autoIncrement bool
	length        int
	array         bool
	unique        bool
	defaultValue  string
}

// ValidCommands :
var validCommands = []string{"--name", "--primary", "--type", "--auto_increment", "--length", "--array", "--unique", "--default"}

// auxillary function to check if item exists in particular array
func in(haystack []string, needle string) bool {
	for _, i := range haystack {
		if needle == i {
			return true
		}
	}
	return false
}

// function to check if given schema is valid as per grammar
func checkIfGrammarCorrect(command string) bool {
	primaryCommand := strings.Split(command, " ")
	for _, value := range primaryCommand {
		if !in(validCommands, strings.Split(value, ":")[0]) {
			return false
		}
	}
	return true
}

func parseAllCommandData(command string) structuredCommandData {
	var data structuredCommandData
	tokens := strings.Split(command, " ")

	for _, value := range tokens {
		var val string
		keyValue := strings.Split(value, ":")
		key := strings.ReplaceAll(keyValue[0], "--", "")
		if len(keyValue) > 1 {
			val = keyValue[1]
		} else {
			val = "true"
		}

		switch key {

		case "name":
			{
				data.name = val
			}
		case "primary":
			{
				data.primary = val == "true"
			}
		case "type":
			{
				data.datatype = val
			}
		case "auto_increment":
			{
				data.autoIncrement = val == "true"
			}
		case "length":
			{
				numericValue, _ := strconv.Atoi(val)
				data.length = numericValue
			}
		case "array":
			{
				data.array = val == "true"
			}
		case "unique":
			{
				data.unique = val == "true"
			}
		case "default":
			{
				data.defaultValue = val
			}
		}
	}
	return data
}

// Process the result from shell to generate a detailed schema
func Process(commands []string) {
	fmt.Println("processing schema, total " + strconv.Itoa(len(commands)))

	// check all input commands for grammar mistakes
	for _, command := range commands {
		if !checkIfGrammarCorrect(command) {
			ui.Error("Invalid grammar : " + command)
		} else {
			ui.Info("parsing all data from command")
			fmt.Println(parseAllCommandData(command))
		}
	}
}

// AutoComplete provides suggestions to user
func AutoComplete(document prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "--name", Description: "a new column in schema"},
		{Text: "--primary", Description: "set column as primary key"},
		{Text: "--auto_increment", Description: "auto increment value of column"},

		{Text: "--type:int", Description: "datatype of column as integer"},
		{Text: "--type:string", Description: "datatype of column as string"},
		{Text: "--type:boolean", Description: "datatype of column as boolean"},
		{Text: "--type:date", Description: "datatype of column as date"},

		{Text: "--length:10", Description: "stores an array"},
		{Text: "--array", Description: "stores an array"},
		{Text: "--unique", Description: "datatype of column as boolean"},
		{Text: "--default:val", Description: "default value of field"},
	}
	return prompt.FilterHasPrefix(suggestions, document.GetWordBeforeCursor(), true)
}
