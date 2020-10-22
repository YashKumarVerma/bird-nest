package generators

import (
	"io/ioutil"
	"strings"

	"github.com/YashKumarVerma/bird-nest/modules/entity"
)

// function to generate destructuring commands
func generateDestructuringString(command entity.StructuredCommandData) string {
	return "\t" + command.Name + ",\n"
}

// function to generate assignment strings
func generateAssignmentString(command entity.StructuredCommandData) string {
	return "\t{{MODULE_NAME_LOWER}}." + command.Name + " = " + command.Name + ";\n"
}

// RepositoryGenerator : function to generate *.repository.ts file
func RepositoryGenerator(moduleName string, schema []entity.StructuredCommandData) (string, string, string) {
	// filenames to access data from
	repositoryTemplateFile := "modules/generators/data/repository.template.js"
	repositoryOutputFile := strings.ToLower(moduleName) + ".repository.ts"

	// reading template content
	templateData, err := ioutil.ReadFile(repositoryTemplateFile)
	Check(err, "error reading "+repositoryTemplateFile)

	// append data to generate final string
	var destructuringString string
	var assignmentString string
	for _, command := range schema {
		destructuringString += generateDestructuringString(command)
		assignmentString += generateAssignmentString(command)
	}
	moduleNameLower := strings.ToLower(moduleName)
	moduleName = strings.Title(moduleNameLower)

	filledTemplate := strings.ReplaceAll(string(templateData), "{{DESTRUCTURING}}", destructuringString)
	filledTemplate = strings.ReplaceAll(filledTemplate, "{{ASSIGNMENT}}", assignmentString)
	filledTemplate = strings.ReplaceAll(filledTemplate, "{{MODULE_NAME}}", moduleName)
	filledTemplate = strings.ReplaceAll(filledTemplate, "{{MODULE_NAME_LOWER}}", moduleNameLower)

	// print the data to console
	// ui.ContextPrint(emoji.Sprint(":bird:"), "finished "+moduleName+".repository.ts")

	// write template file to directory
	return repositoryOutputFile, filledTemplate, moduleName
}
