package generators

import (
	"io/ioutil"
	"strings"

	"github.com/YashKumarVerma/bird-nest/modules/entity"
)

// ServiceGenerator : function to generate *.service.ts file
func ServiceGenerator(moduleName string, schema []entity.StructuredCommandData) (string, string, string) {
	// filenames to access data from
	repositoryTemplateFile := "modules/generators/data/service.template.js"
	repositoryOutputFile := strings.ToLower(moduleName) + ".service.ts"

	// reading template content
	templateData, err := ioutil.ReadFile(repositoryTemplateFile)
	Check(err, "error reading "+repositoryTemplateFile)

	moduleNameLower := strings.ToLower(moduleName)
	moduleName = strings.Title(moduleNameLower)

	filledTemplate := strings.ReplaceAll(string(templateData), "{{MODULE_NAME}}", moduleName)
	filledTemplate = strings.ReplaceAll(filledTemplate, "{{MODULE_NAME_LOWER}}", moduleNameLower)

	// print the data to console
	// ui.ContextPrint(emoji.Sprint(":bird:"), "finished "+moduleName+".service.ts")

	// write template file to directory
	return repositoryOutputFile, filledTemplate, moduleName
}
