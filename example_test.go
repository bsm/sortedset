package sortedset_test

import (
	"fmt"

	"github.com/bsm/sortedset"
)

func ExampleSet() {
	// Create a new set
	set := sortedset.New[string]()

	// Seed with data
	set = set.Add("b")
	set = set.Add("a")
	set = set.Add("c", "a")
	fmt.Println(set.Slice()) // [a b c]

	// Check
	fmt.Println(set.Has("a")) // true
	fmt.Println(set.Has("d")) // false

	// Delete items
	set = set.Delete("a")
	set = set.Delete("d")
	fmt.Println(set.Slice()) // [b c]

	// Output:
	// [a b c]
	// true
	// false
	// [b c]
}
