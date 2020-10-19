package entity

import (
	"fmt"
	"strconv"

	"github.com/c-bata/go-prompt"
)

// AutoComplete provides suggestions to user
func AutoComplete(document prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "column", Description: "a new column in schema"},
		{Text: "--type:string", Description: "datatype of column as string"},
		{Text: "--type:boolean", Description: "datatype of column as boolean"},
		{Text: "--type:integer", Description: "datatype of column as integer"},
		{Text: "--type:date", Description: "datatype of column as date"},
		{Text: "--type:enum", Description: "followed by comma separated enum values"},
		{Text: "--array:true", Description: "stores an array"},
		{Text: "--c:primary", Description: "set column as primary key"},
		{Text: "--c:unique", Description: "datatype of column as boolean"},
		{Text: "--c:default", Description: "default value of field"},
		{Text: "--c:auto_increment", Description: "auto increment value of column"},
		{Text: "--c:null", Description: "allow null values for column"},
	}
	return prompt.FilterHasPrefix(suggestions, document.GetWordBeforeCursor(), true)
}

// Process the result from shell to generate an entity schema
func Process(commands []string) {
	fmt.Println("processing schema, total " + strconv.Itoa(len(commands)))

}
