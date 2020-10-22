package generators

import (
	"io/ioutil"
	"strings"

	"github.com/YashKumarVerma/bird-nest/modules/entity"
)

// ModuleGenerator : function to generate *.service.ts file
func ModuleGenerator(moduleName string, schema []entity.StructuredCommandData) (string, string, string) {
	// filenames to access data from
	serviceTemplateFile := "modules/generators/data/module.template.js"
	serviceOutputFile := strings.ToLower(moduleName) + ".module.ts"

	// reading template content
	templateData, err := ioutil.ReadFile(serviceTemplateFile)
	Check(err, "error reading "+serviceTemplateFile)

	moduleNameLower := strings.ToLower(moduleName)
	moduleName = strings.Title(moduleNameLower)

	filledTemplate := strings.ReplaceAll(string(templateData), "{{MODULE_NAME}}", moduleName)
	filledTemplate = strings.ReplaceAll(filledTemplate, "{{MODULE_NAME_LOWER}}", moduleNameLower)

	// print the data to console
	// ui.ContextPrint(emoji.Sprint(":bird:"), "finished "+moduleName+".module.ts")

	// write template file to directory
	return serviceOutputFile, filledTemplate, moduleName
}
