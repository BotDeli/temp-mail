package api

import (
	"encoding/json"
	"io"
	"net/http"
)

const urlGetDomains = "https://www.1secmail.com/api/v1/?action=getDomainList"

func GetDomains() ([]string, error) {
	response, err := http.Get(urlGetDomains)
	if err != nil {
		return nil, err
	}

	domains, err := decodeJSON(response)
	return domains, err
}

func decodeJSON(response *http.Response) ([]string, error) {
	var content []string

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &content)
	return content, err
}
