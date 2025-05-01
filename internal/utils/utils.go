package utils

import (
	"fmt"
)

func printError(msg string) {
	fmt.Println("❌ [ERROR] " + msg)
	// get_logger().save_log(f"[ERROR] {msg}")
	// raise typer.Exit(1)
}

func printWarning(msg string) {
	// typer.echo(f"⚠️  [WARNING] {msg}")
	// get_logger().save_log(f"[WARNING] {msg}")
}

func printMsg(msg string, enableIcon bool) {
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

func getExtension(filepath string) string {
	// return os.path.splitext(filepath)[1]
	return ""
}

func getDirPath(filepath string) string {
	// return os.path.dirname(filepath)
	return ""
}

func getFileName(filepath string) string {
	// return os.path.splitext(os.path.basename(filepath))[0]
	return ""
}

func createDir(filepath string) {
	// if not os.path.exists(filepath):
	// 	os.makedirs(filepath)
}

func getPath(path string) string {
	// return Path(path)
	return ""
}

func joinPaths(basePath string, subPath string) string {
	// return os.path.join(base_path, sub_path)
	return ""
}

func dirPath(filepath string) string {
	// if os.path.exists(filepath):
	// 	return filepath
	// else:
	// 	error_box("Missing argument '[bold red]FILEPATH[/]'.")
	// 	print_error(f"File '{filepath}' not found.")

	return ""
}
