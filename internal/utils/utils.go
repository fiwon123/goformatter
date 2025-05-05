package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fiwon123/goformatter/internal/logger"
)

func PrintError(msg string) {
	fmt.Println("❌ [ERROR] " + msg)
	logger.SaveLog("[ERROR] " + msg)
	panic("Something goes wrong!")
}

func PrintWarning(msg string) {
	fmt.Println("⚠️ [WARNING] " + msg)
	logger.SaveLog("[WARNING] " + msg)
}

func PrintMsg(msg string, enableIcon bool) {
	if enableIcon {
		fmt.Println("✅ " + msg)
	} else {
		fmt.Println(msg)
	}

	logger.SaveLog(msg)
}

func GetExtension(path string) string {
	_, file := filepath.Split(path)
	fileSplit := strings.Split(file, ".")
	return fileSplit[1]
}

func GetDirPath(path string) string {
	dirPath, _ := filepath.Split(path)
	return dirPath
}

func GetFileName(path string) string {
	_, file := filepath.Split(path)
	fileSplit := strings.Split(file, ".")
	return fileSplit[0]
}

func CreateDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create a single directory
		err := os.MkdirAll(path, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// func GetPath(path string) string {
// 	// return Path(path)
// 	return ""
// }

func JoinPaths(basePath string, subPath string) string {
	return filepath.Join(basePath, subPath)
}

// func DirPath(path string) string {
// 	// if os.path.exists(filepath):
// 	// 	return filepath
// 	// else:
// 	// 	error_box("Missing argument '[bold red]FILEPATH[/]'.")
// 	// 	print_error(f"File '{filepath}' not found.")

// 	return ""
// }
