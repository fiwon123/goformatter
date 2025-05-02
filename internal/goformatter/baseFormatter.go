package goformatter

import (
	"bufio"
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
	check() bool
	dryRun() bool
	format(dirOutput string, inPlace bool) bool
}

func (instance *BaseFormatter) newFormatter(filepath string) {
	instance.FilePath = filepath
}

func (instance *BaseFormatter) serialize(content string) string {
	return ""
}

func (instance *BaseFormatter) deserialize(file string) string {
	return ""
}

func (instance *BaseFormatter) getFormatName() string {
	return instance.ExtensionName
}

func (instance *BaseFormatter) getFormatExtension() string {
	return instance.Extension
}

func (instance *BaseFormatter) check() bool {
	utils.PrintMsg("Checking {self.get_format_name()}...", false)

	file, err := os.Open(instance.FilePath)
	if err != nil {
		utils.PrintError("Invalid " + instance.getFormatName() + ": " + err.Error())
		return false
	}
	defer file.Close()

	contentOriginal := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contentOriginal = scanner.Text()
	}

	content := instance.deserialize(contentOriginal)

	preview := instance.serialize(content)
	if contentOriginal != preview {
		utils.PrintWarning(instance.getFormatName() + " is not correct formatted. But syntax is OK.")
		return false
	}

	utils.PrintMsg(instance.getFormatName()+" is already formatted.", true)

	return true
}

func (instance *BaseFormatter) dryRun() bool {
	utils.PrintMsg(instance.getFormatName()+" Preview:", false)

	file, err := os.Open(instance.FilePath)
	if err != nil {
		utils.PrintError("Invalid " + instance.getFormatName() + ": " + err.Error())
		return false
	}
	defer file.Close()

	contentOriginal := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contentOriginal = scanner.Text()
	}

	content := instance.deserialize(contentOriginal)

	utils.PrintMsg("Deserialized "+instance.getFormatName()+" file.", false)

	// try:
	utils.PrintMsg(instance.serialize(content), false)
	// except TypeError as e:
	// 	print_error(f"Unable to serialize data to {self.get_format_name()}: {e}")

	return true
}

func (instance *BaseFormatter) format(dirOutput string, inPlace bool) bool {
	utils.PrintMsg("Formatting "+instance.getFormatName()+"...", false)

	file, err := os.Open(instance.FilePath)
	if err != nil {
		utils.PrintError("Invalid " + instance.getFormatName() + ": " + err.Error())
		return false
	}
	defer file.Close()

	contentOriginal := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contentOriginal = scanner.Text()
	}

	content := instance.deserialize(contentOriginal)

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

	file, err = os.Open(instance.FilePath)
	if err != nil {
		utils.PrintError("Invalid " + instance.getFormatName() + ": " + err.Error())
		return false
	}
	defer file.Close()

	_, err = file.WriteString(instance.serialize(content))
	if err != nil {
		utils.PrintError("Unable to serialize data to " + instance.getFormatName() + ": " + err.Error())
		return false
	}

	utils.PrintMsg(instance.getFormatName()+" completed.", true)

	return true

}
