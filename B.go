// DO NOT EDIT!  This file was generated via `go generate`

// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy

// B wraps a slice of byte
type B []byte

// Concat returns a copy of the original slice with the additional slice of
// values appended.
func (xs B) Concat(ys []byte) B {
	zs := make(B, len(xs)+len(ys))
	copy(zs, xs)
	copy(zs[len(xs):], ys)
	return zs
}

// ConcatB returns a copy of the original slice with the additional boxed slice
// of values appended.
func (xs B) ConcatB(ys B) B {
	return xs.Concat([]byte(ys))
}

// Contains checks if a value is in the list
func (xs B) Contains(v byte) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

// Elem returns the element with the given index in the list.  Panics if the
// element does not exist.
func (xs B) Elem(n int) byte {
	return xs[n]
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

// Head returns the first value in the list.  Panics if the list is empty.
func (xs B) Head() byte {
	return xs[0]
}

// Init returns a new list with all values except the last.  Panics if the
// list is empty.
func (xs B) Init() B {
	ys := make(B, len(xs)-1)
	copy(ys, xs[:len(xs)-1])
	return ys
}

// Last returns the last value in the list.  Panics if the list is empty.
func (xs B) Last() byte {
	return xs[len(xs)-1]
}

// Len returns the length of the list.
func (xs B) Len() int {
	return len(xs)
}

// Less reports whether the element with index i should sort before the
// element with index j.
func (xs B) Less(i, j int) bool {
	return xs[i] < xs[j]
}

// Map returns a new list with every element transformed by a function
func (xs B) Map(f func(byte) byte) B {
	ys := make(B, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// Reverse returns a copy of the list with the order of elements reversed.
func (xs B) Reverse() B {
	ys := make(B, len(xs))
	n := len(xs) - 1
	for i, v := range xs {
		ys[n-i] = v
	}
	return ys
}

// Swap does an in-place swap of the elements with indexes i and j.  Panics if
// the elements don't exist.  It returns nothing per the Swap signature of
// Sort.Interface.
func (xs B) Swap(i, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}

// Tail returns a new list with all values except the head.  Panics if the
// list is empty.
func (xs B) Tail() B {
	ys := make(B, len(xs)-1)
	copy(ys, xs[1:len(xs)])
	return ys
}

// Unbox returns the list as a native slice
func (xs B) Unbox() []byte {
	return []byte(xs)
}
