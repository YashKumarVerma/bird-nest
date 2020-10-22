package generators

import (
	"io/ioutil"
	"strings"

	"github.com/YashKumarVerma/bird-nest/modules/entity"
	"github.com/YashKumarVerma/bird-nest/modules/ui"
	"github.com/kyokomi/emoji"
)

// ControllerGenerator : function to generate *.controller.ts file
func ControllerGenerator(moduleName string, schema []entity.StructuredCommandData) (string, string, string) {
	// filenames to access data from
	controllerTemplateFile := "modules/generators/data/controller.template.js"
	controllerOutputFile := strings.ToLower(moduleName) + ".controller.ts"

	// reading template content
	templateData, err := ioutil.ReadFile(controllerTemplateFile)
	Check(err, "error reading "+controllerTemplateFile)

	moduleNameLower := strings.ToLower(moduleName)
	moduleName = strings.Title(moduleNameLower)

	filledTemplate := strings.ReplaceAll(string(templateData), "{{MODULE_NAME}}", moduleName)
	filledTemplate = strings.ReplaceAll(filledTemplate, "{{MODULE_NAME_LOWER}}", moduleNameLower)

	// print the data to console
	ui.ContextPrint(emoji.Sprint(":bird:"), "finished "+moduleName+".controller.ts")

	// write template file to directory
	return controllerOutputFile, filledTemplate, moduleName
}
