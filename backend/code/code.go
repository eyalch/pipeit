package code

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

const length = 4

// code represents a code used for pairing-up clients
type code string

// generateCode generates a random 4-digit code consisted of 0-9 digits
func generateCode() code {
	var codeBuilder strings.Builder

	for i := 0; i < length; i++ {
		digit := strconv.Itoa(rand.Intn(10))
		codeBuilder.WriteString(digit)
	}

	return code(codeBuilder.String())
}
