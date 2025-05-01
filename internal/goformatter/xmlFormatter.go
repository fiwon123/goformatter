package goformatter

type XmlFormatter struct {
	BaseFormatter
}

func (instance *XmlFormatter) serialize(content string) {

}

func (instance *XmlFormatter) deserialize(file string) string {
	return ""
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
