package randomizer

import (
	"math/rand"
	"time"
)

type Randomizer interface {
	Intn(n int) int
}

func New() Randomizer {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
