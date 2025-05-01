package utils

import (
	"fmt"
)

func PrintError(msg string) {
	fmt.Println("❌ [ERROR] " + msg)
	// get_logger().save_log(f"[ERROR] {msg}")
	// raise typer.Exit(1)
}

func PrintWarning(msg string) {
	// typer.echo(f"⚠️  [WARNING] {msg}")
	// get_logger().save_log(f"[WARNING] {msg}")
}

func PrintMsg(msg string, enableIcon bool) {
	if enableIcon {
		fmt.Println("✅ " + msg)
	} else {
		fmt.Println(msg)
	}

	// get_logger().save_log(msg)
}

// func errorBox(message string){
// 	console.print(Panel(message, title="[red]Error[/red]", border_style="red"))
// }

func GetExtension(filepath string) string {
	// return os.path.splitext(filepath)[1]
	return ""
}

func GetDirPath(filepath string) string {
	// return os.path.dirname(filepath)
	return ""
}

func GetFileName(filepath string) string {
	// return os.path.splitext(os.path.basename(filepath))[0]
	return ""
}

func CreateDir(filepath string) {
	// if not os.path.exists(filepath):
	// 	os.makedirs(filepath)
}

func GetPath(path string) string {
	// return Path(path)
	return ""
}

func JoinPaths(basePath string, subPath string) string {
	// return os.path.join(base_path, sub_path)
	return ""
}

func DirPath(filepath string) string {
	// if os.path.exists(filepath):
	// 	return filepath
	// else:
	// 	error_box("Missing argument '[bold red]FILEPATH[/]'.")
	// 	print_error(f"File '{filepath}' not found.")

	return ""
}
