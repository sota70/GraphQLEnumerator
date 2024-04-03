package query

import (
	"bytes"
	"fmt"
	"graphqlenumerator/commandargs"
	"graphqlenumerator/jsonbeautifier"
	"io"
	"net/http"
	"strings"

	"github.com/atotto/clipboard"
)

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
