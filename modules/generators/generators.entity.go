package generators

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/YashKumarVerma/bird-nest/modules/entity"
)

// generate javascript code for given struct
func generateNativeJsFromStructForEntity(entity entity.StructuredCommandData) string {
	var js string
	var config string

	// set differnet column type for primary key with auto-increment
	if entity.Primary || entity.AutoIncrement {
		js += "\n@PrimaryGeneratedColumn(CONFIG_STRING)\n"

	} else {
		js += "\n@Column(CONFIG_STRING)\n"

	}

	// set default property
	if entity.DefaultValue != "" {
		config += "default : " + WriteAsPerDataType(entity, entity.DefaultValue) + " ,"
	}

	// set length property
	if entity.Length != 0 {
		config += "length : " + WriteAsPerDataType(entity, strconv.Itoa(entity.Length)) + " ,"
	}

	// set unique property
	if entity.Unique {
		config += "unique : true,"
	}

	// set datatype
	js += entity.Name + " : "
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
	js += ";\n"

	// finalize config string by removing extra commas
	if last := len(config) - 1; last >= 0 && config[last] == ',' {
		config = config[:last]
	}

	configReplaceString := ""
	if config != "" {
		configReplaceString = "{" + config + "}"
	}
	js = strings.ReplaceAll(js, "CONFIG_STRING", configReplaceString)

	return js
}

// EntityGenerator : function to generate *.entity.ts file
func EntityGenerator(moduleName string, schema []entity.StructuredCommandData) (string, string, string) {
	// ui.ContextPrint(emoji.Sprint(":bird:"), "starting "+moduleName+".entity.ts")
	// filenames to access data from
	entityTemplateFile := "modules/generators/data/entity.template.js"
	entityOutputFile := strings.ToLower(moduleName) + ".entity.ts"

	// reading template content
	templateData, err := ioutil.ReadFile(entityTemplateFile)
	Check(err, "error reading "+entityTemplateFile)

	// append data to generate final string
	var codeString string
	for _, command := range schema {
		codeString += generateNativeJsFromStructForEntity(command)
	}
	moduleName = strings.Title(strings.ToLower(moduleName))
	filledTemplate := strings.ReplaceAll(string(templateData), "{{MAIN_WORKSPACE}}", codeString)
	filledTemplate = strings.ReplaceAll(filledTemplate, "{{MODULE_NAME}}", moduleName)

	// print the data to console
	// ui.ContextPrint(emoji.Sprint(":bird:"), "finished "+moduleName+".entity.ts")

	// write template file to directory
	return entityOutputFile, filledTemplate, moduleName
}
