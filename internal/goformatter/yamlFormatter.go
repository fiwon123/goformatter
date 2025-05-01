package goformatter

type YamlFormatter struct {
	BaseFormatter
}

func (instance *YamlFormatter) serialize(content string) {

}

func (instance *YamlFormatter) deserialize(file string) string {
	return ""
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
