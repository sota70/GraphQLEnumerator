package query

import (
	"bytes"
	"fmt"
	"graphqlenumerator/jsonbeautifier"
	"io"
	"net/http"
)

func Query(url string, query string) string {
	if query == "{}" || url == "" {
		return fmt.Sprintf("Usage: ./graphqlenumerator -q -query [query] -u [url]\n")
	}
	var body string = fmt.Sprintf(`{"query": "%s"}`, query)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
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
	return jsonParsedBody
}
