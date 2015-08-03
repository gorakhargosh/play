package shuffle

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// We do not seed the random number generator here to make the PRNG predictable.
// TODO(yesudeep): Bias testing. see:
// http://gregbee.ch/blog/determining-the-bias-of-a-shuffle-algorithm
// TODO(yesudeep): There's a bug below. We aren't testing the possibility of the
// shuffle generating the input sequence without any changes.

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

	// fmt.Println(ss1)
	// fmt.Println(ss2)

	if reflect.DeepEqual(ss1, ss2) {
		t.Fatal("Shuffle on identical string slices failed")
	}

	s3 := make([]string, len(s1))
	copy(s3, s1)
	Strings(s3)
	// fmt.Println(s3)
	if reflect.DeepEqual(s1, s3) {
		t.Fatal("Strings failed")
	}
}

func TestIntSliceShuffle(t *testing.T) {
	s1 := []int{0, 1, 2}
	s2 := make([]int, len(s1))
	copy(s2, s1)

	ss1 := IntSlice(s1)
	ss1.Shuffle()

	ss2 := IntSlice(s2)

	fmt.Println(ss1)
	fmt.Println(ss2)

	if reflect.DeepEqual(ss1, ss2) {
		t.Fatal("Shuffle on identical int slices failed")
	}

	s3 := make([]int, len(s1))
	copy(s3, s1)
	Ints(s3)
	fmt.Println(s3)
	if reflect.DeepEqual(s1, s3) {
		t.Fatal("Ints failed")
	}
}

func TestFloat64SliceShuffle(t *testing.T) {
	s1 := []float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8, 9, 10}
	s2 := make([]float64, len(s1))
	copy(s2, s1)

	ss1 := Float64Slice(s1)
	ss1.Shuffle()

	ss2 := Float64Slice(s2)

	// fmt.Println(ss1)
	// fmt.Println(ss2)

	if reflect.DeepEqual(ss1, ss2) {
		t.Fatal("Shuffle on identical float64 slices failed")
	}

	s3 := make([]float64, len(s1))
	copy(s3, s1)
	Float64s(s3)
	// fmt.Println(s3)
	if reflect.DeepEqual(s1, s3) {
		t.Fatal("Float64s failed")
	}
}
