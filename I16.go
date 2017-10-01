// DO NOT EDIT!  This file was generated via `go generate`

// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy

// I16 wraps a slice of int16
type I16 []int16

// Elem returns the element with the given index in the list.  Panics if the
// element does not exist.
func (xs I16) Elem(n int) int16 {
	return xs[n]
}

// Contains checks if a value is in the list
func (xs I16) Contains(v int16) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

// Filter returns a new list of elements matching a predicate
func (xs I16) Filter(f func(int16) bool) I16 {
	ys := make(I16, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

// Head returns the first value in the list.  Panics if the list is empty.
func (xs I16) Head() int16 {
	return xs[0]
}

// Init returns a new list with all values except the last.  Panics if the
// list is empty.
func (xs I16) Init() I16 {
	ys := make(I16, len(xs)-1)
	copy(ys, xs[:len(xs)-1])
	return ys
}

// List returns the last value in the list.  Panics if the list is empty.
func (xs I16) Last() int16 {
	return xs[len(xs)-1]
}

// Len returns the length of the list.
func (xs I16) Len() int {
	return len(xs)
}

// Less reports whether the element with index i should sort before the
// element with index j.
func (xs I16) Less(i, j int) bool {
	return xs[i] < xs[j]
}

// Map returns a new list with every element transformed by a function
func (xs I16) Map(f func(int16) int16) I16 {
	ys := make(I16, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// Swap does an in-place swap of the elements with indexes i and j.  Panics if
// the elements don't exist.
func (xs I16) Swap(i, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}

// Tail returns a new list with all values except the head.  Panics if the
// list is empty.
func (xs I16) Tail() I16 {
	ys := make(I16, len(xs)-1)
	copy(ys, xs[1:len(xs)])
	return ys
}

// Unbox returns the list as a native slice
func (xs I16) Unbox() []int16 {
	return []int16(xs)
}
