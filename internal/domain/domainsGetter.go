package domain

import (
	"temp-mail/internal/api/gDomains"
	"temp-mail/pkg/randomizer"
	"time"
)

type DomainsGetter struct {
	domains []string
	count   int
	random  randomizer.Randomizer
}

func InitDomainsGetter() *DomainsGetter {
	dom := &DomainsGetter{domains: []string{}, count: 0, random: randomizer.New()}
	go updateDomains(dom)
	time.Sleep(2 * time.Second)
	return dom
}

func updateDomains(dom *DomainsGetter) {
	var (
		newDomains []string
		err        error
	)
	for {
		newDomains, err = gDomains.GetDomains()
		if err != nil {
			time.Sleep(30 * time.Second)
			continue
		}

		dom.domains = newDomains
		dom.count = len(newDomains)
		time.Sleep(5 * time.Minute)
	}
}

func (getter *DomainsGetter) GetRandomDomain() string {
	nextDomain := getter.random.Intn(getter.count)
	return getter.domains[nextDomain]
}
