// DO NOT EDIT!  This file was generated via `go generate`

// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy

// U16 wraps a slice of uint16
type U16 []uint16

// Contains checks if a value is in the list
func (xs U16) Contains(v uint16) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

// Filter returns a new list of elements matching a predicate
func (xs U16) Filter(f func(uint16) bool) U16 {
	ys := make(U16, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

// Map returns a new list with every element transformed by a function
func (xs U16) Map(f func(uint16) uint16) U16 {
	ys := make(U16, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// Unbox returns the list as a native slice
func (xs U16) Unbox() []uint16 {
	return []uint16(xs)
}