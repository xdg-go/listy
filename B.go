// DO NOT EDIT!  This file was generated via `go generate`

// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy

// B wraps a slice of byte
type B []byte

// Contains checks if a value is in the list
func (xs B) Contains(v byte) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

// Filter returns a new list of elements matching a predicate
func (xs B) Filter(f func(byte) bool) B {
	ys := make(B, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

// Map returns a new list with every element transformed by a function
func (xs B) Map(f func(byte) byte) B {
	ys := make(B, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// Unbox returns the list as a native slice
func (xs B) Unbox() []byte {
	return []byte(xs)
}
