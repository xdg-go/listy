// DO NOT EDIT!  This file was generated via `go generate`

// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy

// U16 wraps a slice of uint16
type U16 []uint16

// Concat returns a copy of the original slice with the additional slice of
// values appended.
func (xs U16) Concat(ys []uint16) U16 {
	zs := make(U16, len(xs)+len(ys))
	copy(zs, xs)
	copy(zs[len(xs):], ys)
	return zs
}

// ConcatU16 returns a copy of the original slice with the additional boxed slice
// of values appended.
func (xs U16) ConcatU16(ys U16) U16 {
	return xs.Concat([]uint16(ys))
}

// Contains checks if a value is in the list
func (xs U16) Contains(v uint16) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

// Elem returns the element with the given index in the list.  Panics if the
// element does not exist.
func (xs U16) Elem(n int) uint16 {
	return xs[n]
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

// Head returns the first value in the list.  Panics if the list is empty.
func (xs U16) Head() uint16 {
	return xs[0]
}

// Init returns a new list with all values except the last.  Panics if the
// list is empty.
func (xs U16) Init() U16 {
	ys := make(U16, len(xs)-1)
	copy(ys, xs[:len(xs)-1])
	return ys
}

// Last returns the last value in the list.  Panics if the list is empty.
func (xs U16) Last() uint16 {
	return xs[len(xs)-1]
}

// Len returns the length of the list.
func (xs U16) Len() int {
	return len(xs)
}

// Less reports whether the element with index i should sort before the
// element with index j.
func (xs U16) Less(i, j int) bool {
	return xs[i] < xs[j]
}

// Map returns a new list with every element transformed by a function
func (xs U16) Map(f func(uint16) uint16) U16 {
	ys := make(U16, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// Reverse returns a copy of the list with the order of elements reversed.
func (xs U16) Reverse() U16 {
	ys := make(U16, len(xs))
	n := len(xs) - 1
	for i, v := range xs {
		ys[n-i] = v
	}
	return ys
}

// Swap does an in-place swap of the elements with indexes i and j.  Panics if
// the elements don't exist.  It returns nothing per the Swap signature of
// Sort.Interface.
func (xs U16) Swap(i, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}

// Tail returns a new list with all values except the head.  Panics if the
// list is empty.
func (xs U16) Tail() U16 {
	ys := make(U16, len(xs)-1)
	copy(ys, xs[1:len(xs)])
	return ys
}

// Unbox returns the list as a native slice
func (xs U16) Unbox() []uint16 {
	return []uint16(xs)
}

// Uniq returns a new list with duplicate elements removed.
func (xs U16) Uniq() U16 {
	ys := make(U16, 0)
	seen := make(map[uint16]struct{})
	for _, x := range xs {
		if _, ok := seen[x]; ok {
			continue
		}
		ys = append(ys, x)
		seen[x] = struct{}{}
	}
	return ys
}
