[![GoDoc](https://godoc.org/github.com/xdg/listy?status.svg)](https://godoc.org/github.com/xdg/listy)
[![Build Status](https://travis-ci.org/xdg/listy.svg?branch=master)](https://travis-ci.org/xdg/listy)

# Listy – Functional programming with Golang slices

## Synopsis

```
import (
	"fmt"

	"github.com/xdg/listy"
)

func Example() {

	strings := []string{"foo", "bar", "bar", "baz"}
	ints := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}

	// Wrap slices in listy types to "box" them

	xs := listy.S(strings)
	ys := listy.I(ints)

	// Call methods on boxed types. Chain them as needed.

	if xs.Contains("baz") {
		fmt.Println("Has 'baz'")
	}

	if ys.Contains(9) {
		fmt.Println("Has '9'")
	}

	uxs := xs.Uniq()
	uys := ys.Uniq()

	mxs := uxs.Reverse().Map(func(s string) string { return fmt.Sprintf("'%s'", s) })
	mys := uys.Reverse().Map(func(i int) int { return -1 * i })

	// Unbox them to get built-in slice types back

	fmt.Println("Unique xs:", uxs.Unbox())
	fmt.Println("Unique ys:", uys.Unbox())

	fmt.Println("Reverse mapped unique xs:", mxs.Unbox())
	fmt.Println("Reverse mapped unique ys:", mys.Unbox())

	// Output:
	// Has 'baz'
	// Has '9'
	// Unique xs: [foo bar baz]
	// Unique ys: [3 1 4 5 9 2 6]
	// Reverse mapped unique xs: ['baz' 'bar' 'foo']
	// Reverse mapped unique ys: [-6 -2 -9 -5 -4 -1 -3]
}
```

## Description

This is an experiment-in-progress writing and generating functional list
manipulation functions for Go slices.  The list of functions is inspired
by a subset of the
[Data.List](https://hackage.haskell.org/package/base-4.10.0.0/docs/Data-List.html)
package from Haskell, and similar libraries I've seen in other languages.

The emphasis of this library is on expressiveness and clarity.  With one
exception (`Swap`, for `Sort.Interface`), methods treat inputs as
immutable, returning copies rather than mutating slices in place.

## Copyright and License

Copyright 2017 by David A. Golden. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License"). You may
obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
