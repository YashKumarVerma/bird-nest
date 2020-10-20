package entity

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/YashKumarVerma/bird-nest/modules/ui"
	"github.com/c-bata/go-prompt"
)

// StructuredCommandData to share data to generators
type StructuredCommandData struct {
	Name          string
	Primary       bool
	Datatype      string
	AutoIncrement bool
	Length        int
	Array         bool
	Unique        bool
	DefaultValue  string
}

// ValidCommands : list of all commands that are accepted by shell
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

func parseAllCommandData(command string) StructuredCommandData {
	var data StructuredCommandData
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
				data.Name = val
			}
		case "primary":
			{
				data.Primary = val == "true"
			}
		case "type":
			{
				data.Datatype = val
			}
		case "auto_increment":
			{
				data.AutoIncrement = val == "true"
			}
		case "length":
			{
				numericValue, _ := strconv.Atoi(val)
				data.Length = numericValue
			}
		case "array":
			{
				data.Array = val == "true"
			}
		case "unique":
			{
				data.Unique = val == "true"
			}
		case "default":
			{
				data.DefaultValue = val
			}
		}
	}
	return data
}

// Process the result from shell to generate a detailed schema
func Process(commands []string) []StructuredCommandData {
	fmt.Println("processing schema, total " + strconv.Itoa(len(commands)))

	// expand each schema and generate a structure of all processed data
	entitySchemas := make([]StructuredCommandData, 0)
	for _, command := range commands {
		if !checkIfGrammarCorrect(command) {
			ui.Error("Invalid grammar : " + command)
		} else {
			// ui.Info("parsing all data from command")
			entitySchemas = append(entitySchemas, parseAllCommandData(command))
			parseAllCommandData(command)
		}
	}

	// return array of structures
	return entitySchemas
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
