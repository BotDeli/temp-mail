package generators

import "fmt"

const lengthPattern = 16

func GenerateMailAddress(domain string) string {
	pattern := generateRandomPattern(lengthPattern)
	address := fmt.Sprintf("%s@%s", pattern, domain)
	return address
}
