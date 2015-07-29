package shuffle

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// We do not seed the random number generator here to test reliably
// whether two randomly generated sequences are indeed generated as
// a result of the use of the generator.

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
