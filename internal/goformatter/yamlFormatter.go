package goformatter

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

type YamlFormatter struct {
	BaseFormatter
}

func (instance *YamlFormatter) serialize(content map[string]interface{}) string {
	yamlData, err := yaml.Marshal(content)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(yamlData)
}

func (instance *YamlFormatter) deserialize(file string) map[string]interface{} {
	var result map[string]interface{}

	err := yaml.Unmarshal([]byte(file), &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
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
