// DO NOT EDIT!  This file was generated via `go generate`

// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy

// I wraps a slice of int
type I []int

// Elem returns the element with the given index in the list.  Panics if the
// element does not exist.
func (xs I) Elem(n int) int {
	return xs[n]
}

// Contains checks if a value is in the list
func (xs I) Contains(v int) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

// Filter returns a new list of elements matching a predicate
func (xs I) Filter(f func(int) bool) I {
	ys := make(I, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

// Head returns the first value in the list.  Panics if the list is empty.
func (xs I) Head() int {
	return xs[0]
}

// Init returns a new list with all values except the last.  Panics if the
// list is empty.
func (xs I) Init() I {
	ys := make(I, len(xs)-1)
	copy(ys, xs[:len(xs)-1])
	return ys
}

// List returns the last value in the list.  Panics if the list is empty.
func (xs I) Last() int {
	return xs[len(xs)-1]
}

// Len returns the length of the list.
func (xs I) Len() int {
	return len(xs)
}

// Less reports whether the element with index i should sort before the
// element with index j.
func (xs I) Less(i, j int) bool {
	return xs[i] < xs[j]
}

// Map returns a new list with every element transformed by a function
func (xs I) Map(f func(int) int) I {
	ys := make(I, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// Swap does an in-place swap of the elements with indexes i and j.  Panics if
// the elements don't exist.
func (xs I) Swap(i, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}

// Tail returns a new list with all values except the head.  Panics if the
// list is empty.
func (xs I) Tail() I {
	ys := make(I, len(xs)-1)
	copy(ys, xs[1:len(xs)])
	return ys
}

// Unbox returns the list as a native slice
func (xs I) Unbox() []int {
	return []int(xs)
}
