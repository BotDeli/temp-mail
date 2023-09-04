package generators

import (
	"temp-mail/pkg/randomizer"
)

func NewUUID() string {
	pattern := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	limit := len(pattern)

	random := randomizer.New()

	var UUID string

	for i := 0; i < 16; i++ {
		UUID += string(pattern[random.Intn(limit)])
	}
	return UUID
}
