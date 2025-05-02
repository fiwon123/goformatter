package goformatter

import (
	"encoding/json"
	"fmt"
	"log"
)

type JsonFormatter struct {
	BaseFormatter
}

func (instance *JsonFormatter) serialize(content map[string]interface{}) string {
	jsonData, err := json.Marshal(content)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(jsonData)
}

func (instance *JsonFormatter) deserialize(file string) map[string]interface{} {
	var result map[string]interface{}

	err := json.Unmarshal([]byte(file), &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
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
