package gMail

import (
	"temp-mail/internal/domain"
	"temp-mail/pkg/generators"
)

func GetTempMail(dom *domain.DomainsGetter) (string, error) {
	nextDomain := dom.GetRandomDomain()
	addr := generators.GenerateMailAddress(nextDomain)
	return addr, nil
}
