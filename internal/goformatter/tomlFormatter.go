package goformatter

import (
	"fmt"
	"log"

	"github.com/pelletier/go-toml"
)

type TomlFormatter struct {
	BaseFormatter
}

func (instance *TomlFormatter) serialize(content map[string]interface{}) string {
	tomlData, err := toml.Marshal(content)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(tomlData)
}

func (instance *TomlFormatter) deserialize(file string) map[string]interface{} {
	var result map[string]interface{}

	err := toml.Unmarshal([]byte(file), &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func init() {
	Register("toml", &TomlFormatter{
		BaseFormatter: BaseFormatter{
			FilePath:      "",
			ExtensionName: "toml",
			Extension:     ".toml",
		},
	})
}
