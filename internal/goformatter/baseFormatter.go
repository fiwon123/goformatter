package goformatter

import (
	"os"

	"github.com/fiwon123/goformatter/internal/utils"
)

type BaseFormatter struct {
	FilePath      string
	ExtensionName string
	Extension     string
}

type FormatterInterface interface {
	newFormatter(filepath string)
	serialize(content map[string]interface{}) string
	deserialize(file string) map[string]interface{}
	getFormatName() string
	getFormatExtension() string
	check(f FormatterInterface) bool
	dryRun(f FormatterInterface) bool
	format(f FormatterInterface, dirOutput string, inPlace bool) bool
}

func (instance *BaseFormatter) newFormatter(filepath string) {
	instance.FilePath = filepath
}

func (instance *BaseFormatter) serialize(content map[string]interface{}) string {
	return ""
}

func (instance *BaseFormatter) deserialize(file string) map[string]interface{} {
	return map[string]interface{}{}
}

func (instance *BaseFormatter) getFormatName() string {
	return instance.ExtensionName
}

func (instance *BaseFormatter) getFormatExtension() string {
	return instance.Extension
}

func (instance *BaseFormatter) check(f FormatterInterface) bool {
	utils.PrintMsg("Checking "+instance.getFormatName()+"...", false)

	contentOriginalByte, err := os.ReadFile(instance.FilePath)
	if err != nil {
		utils.PrintError("Invalid " + instance.getFormatName() + ": " + err.Error())
		return false
	}

	contentOriginal := string(contentOriginalByte)

	content := f.deserialize(contentOriginal)

	preview := f.serialize(content)
	if contentOriginal != preview {
		utils.PrintWarning(instance.getFormatName() + " is not correct formatted. But syntax is OK.")
		return false
	}

	utils.PrintMsg(instance.getFormatName()+" is already formatted.", true)

	return true
}

func (instance *BaseFormatter) dryRun(f FormatterInterface) bool {
	utils.PrintMsg(instance.getFormatName()+" Preview:", false)

	contentOriginalByte, err := os.ReadFile(instance.FilePath)
	if err != nil {
		utils.PrintError("Invalid " + instance.getFormatName() + ": " + err.Error())
		return false
	}

	contentOriginal := string(contentOriginalByte)

	content := f.deserialize(contentOriginal)

	utils.PrintMsg("Deserialized "+instance.getFormatName()+" file.", false)

	utils.PrintMsg(f.serialize(content), false)

	return true
}

func (instance *BaseFormatter) format(f FormatterInterface, dirOutput string, inPlace bool) bool {
	utils.PrintMsg("Formatting "+instance.getFormatName()+"...", false)

	contentOriginalByte, err := os.ReadFile(instance.FilePath)
	if err != nil {
		utils.PrintError("Invalid " + instance.getFormatName() + ": " + err.Error())
		return false
	}

	contentOriginal := string(contentOriginalByte)

	println("contentOriginal: " + contentOriginal)

	content := f.deserialize(contentOriginal)

	utils.PrintMsg("Deserialized "+instance.getFormatName()+" file.", false)

	output := ""
	if inPlace {
		output = instance.FilePath
	} else {
		if dirOutput != "" {
			output = dirOutput
			utils.CreateDir(output)
			utils.PrintMsg("Created output folder "+output, false)
		} else {
			output = utils.GetDirPath(instance.FilePath)
			output = utils.JoinPaths(output, "output")
			utils.CreateDir(output)
			utils.PrintMsg("Created output folder "+output, false)
		}

		output = utils.JoinPaths(output, utils.GetFileName(instance.FilePath)+"_formatted"+instance.getFormatExtension())
	}

	err = os.WriteFile(output, []byte(f.serialize(content)), 0644)
	if err != nil {
		utils.PrintError("Unable to serialize data to " + instance.getFormatName() + ": " + err.Error())
		return false
	}

	utils.PrintMsg(instance.getFormatName()+" completed.", true)

	return true

}
