package generators

import (
	"io/ioutil"
	"strings"

	"github.com/YashKumarVerma/bird-nest/modules/entity"
)

// DtoFilterGenerator : function to generate dto/filter-*-dto.ts file
func DtoFilterGenerator(moduleName string, schema []entity.StructuredCommandData) (string, string, string) {
	// filenames to access data from
	templateFile := "./modules/generators/data/dto/filter.dto.js"
	outputFile := "dto/filter-" + strings.ToLower(moduleName) + ".dto.ts"

	// reading template content
	templateData, err := ioutil.ReadFile(templateFile)
	Check(err, "error reading "+outputFile)

	moduleNameLower := strings.ToLower(moduleName)
	moduleName = strings.Title(moduleNameLower)

	filledTemplate := strings.ReplaceAll(string(templateData), "{{MODULE_NAME}}", moduleName)
	filledTemplate = strings.ReplaceAll(filledTemplate, "{{MODULE_NAME_LOWER}}", moduleNameLower)

	// write template file to directory
	return outputFile, filledTemplate, moduleName
}
