package shuffle

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// We do not seed the random number generator here to make the PRNG predictable.

func TestShuffleEmptyCollections(t *testing.T) {
	a := []int{}
	Ints(a)
	if !reflect.DeepEqual(a, []int{}) {
		t.Fatal("Shuffling an empty int slice generates a non-empty slice.")
	}

	b := []string{}
	Strings(b)
	if !reflect.DeepEqual(b, []string{}) {
		t.Fatal("Shuffling an empty string slice generates a non-empty slice.")
	}

	c := []float64{}
	Float64s(c)
	if !reflect.DeepEqual(c, []float64{}) {
		t.Fatal("Shuffling an empty float64 slice generates a non-empty slice.")
	}
}

func TestStringSliceShuffle(t *testing.T) {
	testString := "abcdefghijklmnopqrstuvwxyz"
	s1 := strings.Split(testString, "")
	s2 := make([]string, len(testString))
	copy(s2, s1)

	ss1 := StringSlice(s1)
	ss1.Shuffle()

	ss2 := StringSlice(s2)

	fmt.Println(ss1)
	fmt.Println(ss2)

	if reflect.DeepEqual(ss1, ss2) {
		t.Fatal("Shuffle on identical slices failed")
	}
}

func TestIntSliceShuffle(t *testing.T) {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := make([]int, len(s1))
	copy(s2, s1)

	ss1 := IntSlice(s1)
	ss1.Shuffle()

	ss2 := IntSlice(s2)

	fmt.Println(ss1)
	fmt.Println(ss2)

	if reflect.DeepEqual(ss1, ss2) {
		t.Fatal("Shuffle on identical slices failed")
	}
}
