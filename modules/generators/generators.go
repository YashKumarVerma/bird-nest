package generators

import (
	"os"
	"strconv"

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
func writeDataToDisk(filename string, dataToWrite string) bool {
	// if directory doesn't exist, then create it
	_, err := os.Stat("./generated")
	if os.IsNotExist(err) {
		dirError := os.MkdirAll("./generated", 0755)
		ui.ContextPrint(emoji.Sprint(":package:"), "folder created : generated")
		Check(dirError, "cannot create directory to put fresh content")
	}

	// write data to file
	codeFile, err := os.Create("./generated/" + filename)
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

	ui.ContextPrint(emoji.Sprint(":gear:"), "engine stopping")
}

// Check if error exists, panic
func Check(e error, message string) {
	if e != nil {
		ui.Error(message)
		panic(e)
	}
}
