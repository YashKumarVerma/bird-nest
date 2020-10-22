package generators

import (
	"os"
	"strconv"
	"strings"

	"github.com/YashKumarVerma/bird-nest/modules/entity"
	"github.com/YashKumarVerma/bird-nest/modules/ui"

	"github.com/kyokomi/emoji"
)

// WriteAsPerDataType  to write data as per datatype regulations
func WriteAsPerDataType(entity entity.StructuredCommandData, data string) string {

	if entity.Datatype == "int" || entity.Datatype == "boolean" {
		return data
	}

	return "\"" + data + "\""
}

// WriteDataToDisk to write data to generated directory
func writeDataToDisk(filename string, dataToWrite string, moduleName string) bool {
	moduleName = strings.ToLower(moduleName)
	// ensure that generated directory exists
	_, err := os.Stat("./generated")
	if os.IsNotExist(err) {
		dirError := os.MkdirAll("./generated", 0755)
		Check(dirError, "cannot create directory to put fresh content")
		ui.ContextPrint(emoji.Sprint(":package:"), "folder created : generated")
	}

	// ensure that directory with module name exists
	_, err = os.Stat("./generated/" + moduleName)
	if os.IsNotExist(err) {
		moduleDirErr := os.MkdirAll("./generated/"+moduleName, 0755)
		Check(moduleDirErr, "cannot create directory to put module content")
		ui.ContextPrint(emoji.Sprint(":package:"), "folder created : generated/"+moduleName)
	}

	// ensure that dto directory exists
	_, err = os.Stat("./generated/" + moduleName + "/dto")
	if os.IsNotExist(err) {
		DtoErr := os.MkdirAll("./generated/"+moduleName+"/dto", 0755)
		Check(DtoErr, "cannot create directory to put dto content")
		ui.ContextPrint(emoji.Sprint(":package:"), "folder created : generated/"+moduleName+"/dto")
	}
	// write data to file
	codeFile, err := os.Create("./generated/" + moduleName + "/" + filename)
	Check(err, "error creating "+filename)

	codeBytes, err := codeFile.WriteString(dataToWrite)
	Check(err, "error writing to "+filename)
	ui.ContextPrint(emoji.Sprint(":thread:"), "Total "+strconv.Itoa(codeBytes)+" bytes written to file "+filename)

	fileCloseError := codeFile.Close()
	Check(fileCloseError, "Error closing file "+filename)

	return true
}

// Init to store data locally and call all required methods
func Init(moduleName string, schemas []entity.StructuredCommandData) {
	ui.ContextPrint(emoji.Sprintf(":gear:"), "engine starting")

	// generate .entity.ts file
	writeDataToDisk(EntityGenerator(moduleName, schemas))
	writeDataToDisk(RepositoryGenerator(moduleName, schemas))
	writeDataToDisk(ServiceGenerator(moduleName, schemas))
	writeDataToDisk(ModuleGenerator(moduleName, schemas))
	writeDataToDisk(ControllerGenerator(moduleName, schemas))
	writeDataToDisk(DtoCreateGenerator(moduleName, schemas))
	writeDataToDisk(DtoFilterGenerator(moduleName, schemas))
	writeDataToDisk(DtoUpdateGenerator(moduleName, schemas))

	ui.ContextPrint(emoji.Sprint(":gear:"), "engine stopping")
	ui.ContextPrint(emoji.Sprint(":construction:"), "search for @NeedsManualIntervention for all places where manual edits are needed")
}

// Check if error exists, panic
func Check(e error, message string) {
	if e != nil {
		ui.Error(message)
		panic(e)
	}
}
