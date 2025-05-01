package goformatter

type BaseFormatter struct {
	FilePath      string
	ExtensionName string
	Extension     string
}

type FormatterInterface interface {
	newFormatter(filepath string)
	serialize(content string)
	deserialize(file string) string
	getFormatName() string
	getFormatExtension() string
	check() bool
	dryRun() bool
	format(dirOutput string, inPlace bool) bool
}

func (instance *BaseFormatter) newFormatter(filepath string) {
	instance.FilePath = filepath
}

func (instance *BaseFormatter) serialize(content string) {

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
	// print_msg(f"Checking {self.get_format_name()}...")

	// try:
	// 	with open(self.filepath, "r", encoding="utf-8") as file:
	// 		content_original = file.read()
	// 	with open(self.filepath, "r", encoding="utf-8") as file:
	// 		content = self._deserialize(file)
	// except Exception as e:
	// 	print_error(f"Invalid {self.get_format_name()}: {e}")

	// default_comma = ", "
	// default_colon = " : "

	// preview = self._serialize(content, indent=4, separators=(default_comma, default_colon))
	// if (content_original != preview):
	// 	print_warning(f"{self.get_format_name()} is not correct formatted. But syntax is OK.")
	// 	return False

	// print_msg(f"{self.get_format_name()} is already formatted.", True)

	return true
}

func (instance *BaseFormatter) dryRun() bool {
	// print_msg(f"{self.get_format_name()} Preview:")

	// try:
	// 	with open(self.filepath, "r", encoding="utf-8") as file:
	// 		content = self._deserialize(file)
	// except Exception as e:
	// 	print_error(f"Invalid {self.get_format_name()}: {e}")

	// print_msg(f"Deserialized {self.get_format_name()} file.")

	// default_comma = ", "
	// default_colon = " : "

	// try:
	// 	print_msg(self._serialize(content, indent=4, separators=(default_comma, default_colon)))
	// except TypeError as e:
	// 	print_error(f"Unable to serialize data to {self.get_format_name()}: {e}")

	return true
}

func (instanec *BaseFormatter) format(dirOutput string, inPlace bool) bool {
	// print_msg(f"Formatting {self.get_format_name()}...")

	// try:
	// 	with open(self.filepath, "r", encoding="utf-8") as file:
	// 		content = self._deserialize(file)
	// except Exception as e:
	// 	print_error(f"Invalid {self.get_format_name()}: {e}")

	// print_msg(f"Deserialized {self.get_format_name()} file.")
	// default_comma = ", "
	// default_colon = " : "

	// if (in_place):
	// 	output = self.filepath
	// else:
	// 	if (dir_output != None):
	// 		output = dir_output
	// 		create_dir(output)
	// 		print_msg(f"Created output folder {output}")
	// 	else:
	// 		output = get_dir_path(self.filepath)
	// 		output = join_paths(output, "output")
	// 		create_dir(output)
	// 		print_msg(f"Created output folder {output}")

	// 	output = join_paths(output, get_file_name(self.filepath) + "_formatted" + self.get_format_extension())

	// try:
	// 	with open(output, "w", encoding="utf-8") as file:
	// 		file.write(self._serialize(content, indent=4, separators=(default_comma, default_colon)))
	// except TypeError as e:
	// 	print_error(f"Unable to serialize data to {self.get_format_name()}: {e}")

	// print_msg(f"{self.get_format_name()} completed.", True)

	return true

}
