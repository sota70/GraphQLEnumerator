/*
Query package

includes a function that executes GraphQL query
*/
package query

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"example.com/graphqlenumerator/commandargs"
	"example.com/graphqlenumerator/jsonbeautifier"

	"github.com/atotto/clipboard"
)

/*
Query function

# Overview

executes GraphQL query

usage is following:

	./graphqlenumerator -q -query [query] -u [url(graphql endpoint)]

(!) Command arguments must indicate that is is on query mode and include query and graphql

# Parameters

args commandargs.CommandArgs:

	Command arguments

# Return

returns JSON parsed HTTP response
*/
func Query(args commandargs.CommandArgs) string {
	if *args.Query == "{}" || *args.U == "" {
		return fmt.Sprintf("Usage: ./graphqlenumerator -q -query [query] -u [url]\n")
	}
	// query cannot have new-line character
	var query string = strings.ReplaceAll(*args.Query, "\n", "")
	var body string = fmt.Sprintf(`{"query": "%s"}`, query)
	request, err := http.NewRequest("POST", *args.U, bytes.NewBuffer([]byte(body)))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return fmt.Sprintf("Error During Query: %v\n", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error During Query: %v\n", err)
	}
	jsonParsedBody, err := jsonbeautifier.BeautifyJSON(string(respBody), 2)
	if err != nil {
		return fmt.Sprintf("Error During Parsing JSON: %v\n", err)
	}
	if *args.C {
		clipboard.WriteAll(jsonParsedBody)
	}
	return jsonParsedBody
}
