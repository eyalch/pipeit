package code

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	seed   = 1
	first  = "1779"
	second = "1850"
	third  = "6041"
)

func TestGenerateCode(t *testing.T) {
	rand.Seed(seed)

	tests := []code{first, second, third}

	for i, want := range tests {
		t.Run(fmt.Sprintf("code%d", i), func(t *testing.T) {
			got := generateCode()

			if got != want {
				t.Errorf("got %s; want %s", got, want)
			}
		})
	}
}
