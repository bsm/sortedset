# Sorted Set

[![Go Reference](https://pkg.go.dev/badge/github.com/bsm/sortedset.svg)](https://pkg.go.dev/github.com/bsm/sortedset)
[![Test](https://github.com/bsm/sortedset/actions/workflows/test.yml/badge.svg)](https://github.com/bsm/sortedset/actions/workflows/test.yml)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Simple set implementation. Uses sorted slices and generics.

### Documentation

Full documentation is available on [GoDoc](http://godoc.org/github.com/bsm/sortedset)

### Example

```go
package main

import (
  "fmt"

  "github.com/bsm/sortedset"
)

func main() {
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
}
```
