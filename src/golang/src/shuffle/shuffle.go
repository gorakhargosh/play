// Package shuffle implements a Fisher-Yates (or Knuth) shuffle for shuffling
// slices and user-defined collections.
//
// In order to guarantee a successful random shuffle, the pseudo-random number
// generator must be correctly seeded before using any of the functions defined
// in this package.
//
// See: https://en.wikipedia.org/wiki/Fisher-Yates_shuffle
package shuffle

import "math/rand"

// A type, typically a collection, that satisfies shuffle.Interface can be
// shuffled by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// Fisher-Yates shuffle, or Knuth shuffle, that shuffles an indexable
// collection of items.
func Shuffle(data Interface) {
	for n := data.Len(); n > 0; {
		i := rand.Intn(n)
		n--
		data.Swap(n, i)
	}
}

// Convenience type for most common cases.
type IntSlice []int

func (s IntSlice) Len() int      { return len(s) }
func (s IntSlice) Swap(i, j int) { t := s[i]; s[i] = s[j]; s[j] = t }

// Shuffle is a convenience method.
func (s IntSlice) Shuffle() { Shuffle(s) }

type StringSlice []string

func (s StringSlice) Len() int      { return len(s) }
func (s StringSlice) Swap(i, j int) { t := s[i]; s[i] = s[j]; s[j] = t }

// Shuffle is a convenience method.
func (s StringSlice) Shuffle() { Shuffle(s) }

type Float64Slice []float64

func (s Float64Slice) Len() int      { return len(s) }
func (s Float64Slice) Swap(i, j int) { t := s[i]; s[i] = s[j]; s[j] = t }

// Shuffle is a convenience method.
func (s Float64Slice) Shuffle() { Shuffle(s) }

// Convenience wrappers.
func Ints(a []int)         { Shuffle(IntSlice(a)) }
func Strings(a []string)   { Shuffle(StringSlice(a)) }
func Float64s(a []float64) { Shuffle(Float64Slice(a)) }
