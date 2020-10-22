package generators

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/YashKumarVerma/bird-nest/modules/entity"
)

// generate javascript code for given struct
func generateNativeJsFromStructForDtoUpdate(entity entity.StructuredCommandData) string {
	var js string

	if entity.Name == "id" {
		js += "\t@ApiProperty({\n"
		js += "\t\tdescription: 'something',\n"
		js += "\t\trequired: true\n"
		js += "\t})\n"
		js += "\t@IsNumber()\n"
		js += "\tid: number;\n\n"
		return js
	}

	js += "\t@ApiProperty({\n"
	js += "\t\tdescription: 'something',\n"
	js += "\t\trequired: false\n"
	js += "\t})\n"
	js += "\t@IsOptional()\n"

	// set length property
	if entity.Length != 0 {
		js += "\t@Length(" + strconv.Itoa(entity.Length) + ")\n"
	}

	// set datatype
	js += "\t" + entity.Name + " : "
	switch entity.Datatype {
	case "int":
		{
			js += "number"
		}
	case "string":
		{
			js += "string"
		}
	case "date":
		{
			js += "string"
		}
	case "boolean":
		{
			js += "boolean"

		}
	}
	js += ";\n\n"

	return js
}

// DtoUpdateGenerator : function to generate dto/update-*-dto.ts file
func DtoUpdateGenerator(moduleName string, schema []entity.StructuredCommandData) (string, string, string) {
	// filenames to access data from
	templateFile := "./modules/generators/data/dto/update.dto.js"
	outputFile := "dto/update-" + strings.ToLower(moduleName) + ".dto.ts"

	// reading template content
	templateData, err := ioutil.ReadFile(templateFile)
	Check(err, "error reading "+outputFile)

	// append data to generate final string
	var codeString string
	for _, command := range schema {
		codeString += generateNativeJsFromStructForDtoUpdate(command)
	}

	moduleNameLower := strings.ToLower(moduleName)
	moduleName = strings.Title(moduleNameLower)

	filledTemplate := strings.ReplaceAll(string(templateData), "{{MAIN_WORKSPACE}}", codeString)
	filledTemplate = strings.ReplaceAll(filledTemplate, "{{MODULE_NAME}}", moduleName)
	filledTemplate = strings.ReplaceAll(filledTemplate, "{{MODULE_NAME_LOWER}}", moduleNameLower)

	// write template file to directory
	return outputFile, filledTemplate, moduleName
}
