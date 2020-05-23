package code

import (
	"math/rand"
	"reflect"
	"testing"
)

func assertDeepEqual(t *testing.T, got, want []code) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func setUp() store {
	rand.Seed(seed)

	var s store
	s.create() // first
	s.create() // second
	s.create() // third

	return s
}

func TestStore_create(t *testing.T) {
	s := setUp()

	assertDeepEqual(t, s.codes, []code{first, second, third})
}

func TestStore_createDuplicate(t *testing.T) {
	s := setUp()

	// Reset the seed
	rand.Seed(1)

	got := s.create()

	// Assert that creating a new code generates a code different from the first one,
	// since resetting the seed should produce the same order of random codes.
	if got == first {
		t.Errorf("the new code should not be equal the first one (%s)", first)
	}
}

func TestStore_delete(t *testing.T) {
	s := setUp()

	s.delete(second)

	assertDeepEqual(t, s.codes, []code{first, third})
}

func TestStore_deleteNotExists(t *testing.T) {
	s := setUp()

	s.delete("1234")

	assertDeepEqual(t, s.codes, []code{first, second, third})
}
