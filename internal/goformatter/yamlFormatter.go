package goformatter

type YamlFormatter struct {
	BaseFormatter
}

func (instance *YamlFormatter) serialize(content string) string {
	return ""
}

func (instance *YamlFormatter) deserialize(file string) map[string]interface{} {
	return map[string]interface{}{}
}

func init() {
	Register(".yaml", &YamlFormatter{
		BaseFormatter: BaseFormatter{
			FilePath:      "",
			ExtensionName: "yaml",
			Extension:     ".yaml",
		},
	})
}
