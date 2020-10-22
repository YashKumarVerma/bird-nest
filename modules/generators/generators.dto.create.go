package generators

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/YashKumarVerma/bird-nest/modules/entity"
)

// generate javascript code for given struct
func generateNativeJsFromStructForDtoCreate(entity entity.StructuredCommandData) string {
	var js string = "\t@ApiProperty({ description : 'something'})\n"

	// set length property
	if entity.Length != 0 {
		js += "\t@Length(" + strconv.Itoa(entity.Length) + ")\n"
	}

	// check if value required to create new entity
	if entity.Null == false {
		js += "\t@IsNotEmpty()\n"
	}

	if entity.Datatype == "string" {
		js += "\t@IsAlpha()\n"
	}

	// put all logic for code inclusion above this
	js += "\t" + entity.Name + ": " + entity.Datatype + ";\n\n"
	return js
}

// DtoCreateGenerator : function to generate dto/create-*-dto.ts file
func DtoCreateGenerator(moduleName string, schema []entity.StructuredCommandData) (string, string, string) {
	// filenames to access data from
	templateFile := "./modules/generators/data/dto/create.dto.js"
	outputFile := "dto/create-" + strings.ToLower(moduleName) + ".dto.ts"

	// reading template content
	templateData, err := ioutil.ReadFile(templateFile)
	Check(err, "error reading "+outputFile)

	// append data to generate final string
	var codeString string
	for _, command := range schema {
		codeString += generateNativeJsFromStructForDtoCreate(command)
	}

	moduleNameLower := strings.ToLower(moduleName)
	moduleName = strings.Title(moduleNameLower)

	filledTemplate := strings.ReplaceAll(string(templateData), "{{MAIN_WORKSPACE}}", codeString)
	filledTemplate = strings.ReplaceAll(filledTemplate, "{{MODULE_NAME}}", moduleName)
	filledTemplate = strings.ReplaceAll(filledTemplate, "{{MODULE_NAME_LOWER}}", moduleNameLower)

	// write template file to directory
	return outputFile, filledTemplate, moduleName
}
