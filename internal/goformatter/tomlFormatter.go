package goformatter

type TomlFormatter struct {
	BaseFormatter
}

func (instance *TomlFormatter) serialize(content string) string {
	return ""
}

func (instance *TomlFormatter) deserialize(file string) map[string]interface{} {
	return map[string]interface{}{}
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
