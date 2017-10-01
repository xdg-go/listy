// DO NOT EDIT!  This file was generated via `go generate`

// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy

// U8 wraps a slice of uint8
type U8 []uint8

// Contains checks if a value is in the list
func (xs U8) Contains(v uint8) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

// Filter returns a new list of elements matching a predicate
func (xs U8) Filter(f func(uint8) bool) U8 {
	ys := make(U8, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

// Head returns the first value in the list.  Panics if the list is empty.
func (xs U8) Head() uint8 {
	return xs[0]
}

// Init returns a new list with all values except the last.  Panics if the
// list is empty.
func (xs U8) Init() U8 {
	ys := make(U8, len(xs)-1)
	copy(ys, xs[:len(xs)-1])
	return ys
}

// List returns the last value in the list.  Panics if the list is empty.
func (xs U8) Last() uint8 {
	return xs[len(xs)-1]
}

// Len returns the length of the list.
func (xs U8) Len() int {
	return len(xs)
}

// Less reports whether the element with index i should sort before the
// element with index j.
func (xs U8) Less(i, j int) bool {
	return xs[i] < xs[j]
}

// Map returns a new list with every element transformed by a function
func (xs U8) Map(f func(uint8) uint8) U8 {
	ys := make(U8, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// Swap swaps the elements with indexes i and j and returns the original list.
func (xs U8) Swap(i, j int) U8 {
	xs[i], xs[j] = xs[j], xs[i]
	return xs
}

// Tail returns a new list with all values except the head.  Panics if the
// list is empty.
func (xs U8) Tail() U8 {
	ys := make(U8, len(xs)-1)
	copy(ys, xs[1:len(xs)])
	return ys
}

// Unbox returns the list as a native slice
func (xs U8) Unbox() []uint8 {
	return []uint8(xs)
}
