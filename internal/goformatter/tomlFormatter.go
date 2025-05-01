package goformatter

type TomlFormatter struct {
	BaseFormatter
}

func (instance *TomlFormatter) serialize(content string) {

}

func (instance *TomlFormatter) deserialize(file string) string {
	return ""
}

func init() {
	Register(".toml", &TomlFormatter{
		BaseFormatter: BaseFormatter{
			FilePath:      "",
			ExtensionName: "toml",
			Extension:     ".toml",
		},
	})
}
