package gDomains

import (
	"net/http"
	"temp-mail/pkg/decoder"
)

const urlGetDomains = "https://www.1secmail.com/api/v1/?action=getDomainList"

func GetDomains() ([]string, error) {
	response, err := http.Get(urlGetDomains)
	if err != nil {
		return nil, err
	}

	domains, err := decoder.DecodeJSONToArray[string](response)
	return domains, err
}
