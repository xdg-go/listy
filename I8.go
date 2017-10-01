// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy

// I8 wraps a slice of int8
type I8 []int8

// Elem returns the element with the given index in the list.  Panics if the
// element does not exist.
func (xs I8) Elem(n int) int8 {
	return xs[n]
}

// Contains checks if a value is in the list
func (xs I8) Contains(v int8) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

// Filter returns a new list of elements matching a predicate
func (xs I8) Filter(f func(int8) bool) I8 {
	ys := make(I8, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

// Head returns the first value in the list.  Panics if the list is empty.
func (xs I8) Head() int8 {
	return xs[0]
}

// Init returns a new list with all values except the last.  Panics if the
// list is empty.
func (xs I8) Init() I8 {
	ys := make(I8, len(xs)-1)
	copy(ys, xs[:len(xs)-1])
	return ys
}

// List returns the last value in the list.  Panics if the list is empty.
func (xs I8) Last() int8 {
	return xs[len(xs)-1]
}

// Len returns the length of the list.
func (xs I8) Len() int {
	return len(xs)
}

// Less reports whether the element with index i should sort before the
// element with index j.
func (xs I8) Less(i, j int) bool {
	return xs[i] < xs[j]
}

// Map returns a new list with every element transformed by a function
func (xs I8) Map(f func(int8) int8) I8 {
	ys := make(I8, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// Swap does an in-place swap of the elements with indexes i and j.  Panics if
// the elements don't exist.
func (xs I8) Swap(i, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}

// Tail returns a new list with all values except the head.  Panics if the
// list is empty.
func (xs I8) Tail() I8 {
	ys := make(I8, len(xs)-1)
	copy(ys, xs[1:len(xs)])
	return ys
}

// Unbox returns the list as a native slice
func (xs I8) Unbox() []int8 {
	return []int8(xs)
}
