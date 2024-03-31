package jsonbeautifier

import (
	"bytes"
	"encoding/json"
)

func BeautifyJSON(jsonStr string, indent int) (string, error) {
	var beautifiedJSON bytes.Buffer
	var indentStr string = ""
	for i := 0; i < indent; i++ {
		indentStr += " "
	}
	if err := json.Indent(&beautifiedJSON, []byte(jsonStr), "", indentStr); err != nil {
		return "", err
	}
	return beautifiedJSON.String(), nil
}
