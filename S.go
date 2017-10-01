// DO NOT EDIT!  This file was generated via `go generate`

// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy

// S wraps a slice of string
type S []string

// Elem returns the element with the given index in the list.  Panics if the
// element does not exist.
func (xs S) Elem(n int) string {
	return xs[n]
}

// Contains checks if a value is in the list
func (xs S) Contains(v string) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

// Filter returns a new list of elements matching a predicate
func (xs S) Filter(f func(string) bool) S {
	ys := make(S, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

// Head returns the first value in the list.  Panics if the list is empty.
func (xs S) Head() string {
	return xs[0]
}

// Init returns a new list with all values except the last.  Panics if the
// list is empty.
func (xs S) Init() S {
	ys := make(S, len(xs)-1)
	copy(ys, xs[:len(xs)-1])
	return ys
}

// List returns the last value in the list.  Panics if the list is empty.
func (xs S) Last() string {
	return xs[len(xs)-1]
}

// Len returns the length of the list.
func (xs S) Len() int {
	return len(xs)
}

// Less reports whether the element with index i should sort before the
// element with index j.
func (xs S) Less(i, j int) bool {
	return xs[i] < xs[j]
}

// Map returns a new list with every element transformed by a function
func (xs S) Map(f func(string) string) S {
	ys := make(S, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// Swap does an in-place swap of the elements with indexes i and j.  Panics if
// the elements don't exist.
func (xs S) Swap(i, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}

// Tail returns a new list with all values except the head.  Panics if the
// list is empty.
func (xs S) Tail() S {
	ys := make(S, len(xs)-1)
	copy(ys, xs[1:len(xs)])
	return ys
}

// Unbox returns the list as a native slice
func (xs S) Unbox() []string {
	return []string(xs)
}
