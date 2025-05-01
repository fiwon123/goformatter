package goformatter

type XmlFormatter struct {
	BaseFormatter
}

func (instance *XmlFormatter) serialize(content string) string {
	return ""
}

func (instance *XmlFormatter) deserialize(file string) map[string]interface{} {
	return map[string]interface{}{}
}

func init() {
	Register(".xml", &XmlFormatter{
		BaseFormatter: BaseFormatter{
			FilePath:      "",
			ExtensionName: "xml",
			Extension:     ".xml",
		},
	})
}
