/*
Json beautifier package

beautifies JSON
*/
package jsonbeautifier

import (
	"bytes"
	"encoding/json"
)

/*
BeautifyJSON function

# Overview

beautifies JSON

makes an indent

# Parameter

jsonStr string:

	old json string

indent int:

	how many spaces in one indent

# Return

returns a tuple of a new json string and an error (if there is an error)
*/
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
