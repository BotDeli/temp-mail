package generators

import (
	"temp-mail/pkg/randomizer"
)

const (
	chars       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	lengthChars = len(chars)
)

func generateRandomPattern(length int) string {
	var (
		bytes = make([]byte, length)
		num   int
	)

	random := randomizer.New()

	for i := 0; i < length; i++ {
		num = random.Intn(lengthChars)
		bytes[i] = chars[num]
	}
	return string(bytes)
}
