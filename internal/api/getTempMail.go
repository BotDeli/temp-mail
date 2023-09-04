package api

import (
	"temp-mail/internal/generators"
)

func GetTempMail() (string, error) {
	domains, err := GetDomains()
	if err != nil {
		return "", err
	}

	addr := generators.GenerateMailAddress(domains[0])
	return addr, nil
}
