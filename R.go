// DO NOT EDIT!  This file was generated via `go generate`

// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy

// R wraps a slice of rune
type R []rune

// Contains checks if a value is in the list
func (xs R) Contains(v rune) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

// Filter returns a new list of elements matching a predicate
func (xs R) Filter(f func(rune) bool) R {
	ys := make(R, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

// Map returns a new list with every element transformed by a function
func (xs R) Map(f func(rune) rune) R {
	ys := make(R, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// Unbox returns the list as a native slice
func (xs R) Unbox() []rune {
	return []rune(xs)
}
