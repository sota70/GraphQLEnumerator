package enumeration

import (
	"bytes"
	"fmt"
	"graphqlenumerator/jsonbeautifier"
	"io"
	"net/http"
)

func Enumerate(url string) string {
	if url == "" {
		return fmt.Sprintf("Usage: ./graphqlenumerator [graphql endpoint url]\n")
	}
	var values string = `{"query": "query {__schema {types {name,fields {name,type { name}}}}}"}`
	req, err := http.NewRequest(
		"POST",
		url,
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
	return jsonParsedBody
}
