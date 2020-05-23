package code

// store is responsible for holding the active codes as well as creating and revoking codes
type store struct {
	codes []code
}

func (s *store) findIndex(code code) int {
	index := -1
	for i, c := range s.codes {
		if c == code {
			index = i
			break
		}
	}
	return index
}

// create generates and registers a random code.
func (s *store) create() code {
	code := generateCode()

	// Keep generating new codes until a new one is generated
	for s.findIndex(code) != -1 {
		code = generateCode()
	}

	s.codes = append(s.codes, code)
	return code
}

// delete removes the given code from the codes slice
func (s *store) delete(code code) {
	index := s.findIndex(code)

	// If the given code doesn't exist, do nothing
	if index == -1 {
		return
	}

	// Remove the given code from the codes slice
	s.codes[index] = s.codes[len(s.codes)-1]
	s.codes = s.codes[:len(s.codes)-1]
}
