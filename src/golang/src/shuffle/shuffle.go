// Package shuffle implements a Fisher-Yates (or Knuth) shuffle for a
// collection that satisfies the interface defined in this package.
//
// In order to guarantee a successful and random shuffle, the pseudo-random
// number generator must be correctly seeded before using any of the functions
// defined in this package.
package shuffle

import "math/rand"

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// Fisher-Yates shuffle, or Knuth shuffle, that shuffles an indexable
// collection of items.
func Shuffle(col Interface) {
	for n := col.Len(); n > 0; {
		i := rand.Intn(n)
		n--
		col.Swap(n, i)
	}
}

// A slice of string.
type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Swap(i, j int) {
	t := s[i]
	s[i] = s[j]
	s[j] = t
}

// Convenience method.
func (s StringSlice) Shuffle() {
	Shuffle(s)
}
