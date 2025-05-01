package goformatter

type JsonFormatter struct {
	BaseFormatter
}

func (instance *JsonFormatter) serialize(content string) {

}

func (instance *JsonFormatter) deserialize(file string) string {
	return ""
}

func init() {
	Register(".json", &JsonFormatter{
		BaseFormatter: BaseFormatter{
			FilePath:      "",
			ExtensionName: "json",
			Extension:     ".json",
		},
	})
}
