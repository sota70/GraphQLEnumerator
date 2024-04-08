/*
Enumeration package

enumerates GraphQL schema abusing introspection
*/
package enumeration

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"example.com/graphqlenumerator/commandargs"
	"example.com/graphqlenumerator/jsonbeautifier"

	"github.com/atotto/clipboard"
)

/*
Enumerate function

# Overview

enumerates GraphQL schema abusing introspection

(!) command arguments must indicate that it is on enumerate mode and contain GraphQL endpoint URL

usage is following:

	./graphqlenumerator -e -u [graphql endpoint url]

# Parameters

args commandargs.CommandArgs:

	command arguments

# Return

returns JSON parsed HTTP response
*/
func Enumerate(args commandargs.CommandArgs) string {
	if *args.U == "" {
		return fmt.Sprintf("Usage: ./graphqlenumerator -e -u [graphql endpoint url]\n")
	}
	var values string = `{"query": "query {__schema {types {name,fields {name,type { name}}}}}"}`
	req, err := http.NewRequest(
		"POST",
		*args.U,
		bytes.NewBuffer([]byte(values)),
	)
	if err != nil {
		return fmt.Sprintf("Error During POST Request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("Error Dugin POST Request: %v\n", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error During POST Request: %v\n", err)
	}
	jsonParsedBody, err := jsonbeautifier.BeautifyJSON(string(body), 2)
	if err != nil {
		return fmt.Sprintf("Error During Parsing JSON: %v\n", err)
	}
	if *args.C {
		clipboard.WriteAll(jsonParsedBody)
	}
	return jsonParsedBody
}
