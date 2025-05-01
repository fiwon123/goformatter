package goformatter

import (
	"github.com/fiwon123/goformatter/internal/utils"
)

func Process(filepath string, dirOutput string, isChecking bool, isDryRun bool, isPretty bool, inPlace bool, fileType string) {
	utils.PrintMsg("-------------------------------------", false)
	utils.PrintMsg("Verifying filepath...", false)
	var formatterClass FormatterInterface
	ok := false
	if filepath != "" {
		formatterClass, ok = GetFormatter(filepath)

		if !ok {
			utils.PrintError("Invalid parameter file_type: " + fileType)
		}
	} else {
		formatterClass, ok = GetFormatter(utils.GetExtension(filepath)[1:])

		if !ok {
			utils.PrintError("Invalid input file: " + filepath)
		}
	}

	utils.PrintMsg("Verification passed.", false)
	formatterClass.newFormatter(filepath)

	if isChecking {
		formatterClass.check()
	}
	if isDryRun {
		formatterClass.dryRun()
	}
	if !isChecking && !isDryRun {
		formatterClass.format(dirOutput, inPlace)
	}
}
