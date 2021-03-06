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
// elements of the collection be enumerated by an integer index. If you want to
// use the shuffling algorithm with a random number generator of your choice,
// you can choose to implement RandIntn differently from what we have
// implemented.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)

	// Generate a random number in [0, n).
	RandIntn(n int) int
}

// shuffle ensures that the shuffled subarray is toward the start of the array,
// not the right end. This allows us to shuffle only count number of items in
// the array, if required.
func shuffle(data Interface, count int) {
	len := data.Len()
	if count > len {
		count = len
	}
	var k = 0
	for i := 0; i < count; i++ {
		k = data.RandIntn(len-i) + i // Scale the random number.
		data.Swap(k, i)
	}
}

// Shuffle performs a Fisher-Yates shuffle, or Knuth shuffle, on an indexable
// collection of items.
func Shuffle(data Interface) {
	shuffle(data, data.Len())

	// Simple right-subarray implementation.
	//
	// for n := data.Len(); n > 0; {
	// 	i := data.RandIntn(n)
	// 	n--
	// 	data.Swap(n, i)
	// }
}

// Convenience type for most common cases.
type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { t := s[i]; s[i] = s[j]; s[j] = t }
func (s IntSlice) RandIntn(n int) int { return rand.Intn(n) }

// Shuffle is a convenience method.
func (s IntSlice) Shuffle() { Shuffle(s) }

type StringSlice []string

func (s StringSlice) Len() int           { return len(s) }
func (s StringSlice) Swap(i, j int)      { t := s[i]; s[i] = s[j]; s[j] = t }
func (s StringSlice) RandIntn(n int) int { return rand.Intn(n) }

// Shuffle is a convenience method.
func (s StringSlice) Shuffle() { Shuffle(s) }

type Float64Slice []float64

func (s Float64Slice) Len() int           { return len(s) }
func (s Float64Slice) Swap(i, j int)      { t := s[i]; s[i] = s[j]; s[j] = t }
func (s Float64Slice) RandIntn(n int) int { return rand.Intn(n) }

// Shuffle is a convenience method.
func (s Float64Slice) Shuffle() { Shuffle(s) }

// Convenience wrappers.
func Ints(a []int)         { Shuffle(IntSlice(a)) }
func Strings(a []string)   { Shuffle(StringSlice(a)) }
func Float64s(a []float64) { Shuffle(Float64Slice(a)) }
