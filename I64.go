// DO NOT EDIT!  This file was generated via `go generate`

// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy

// I64 wraps a slice of int64
type I64 []int64

// Concat returns a copy of the original slice with the additional slice of
// values appended.
func (xs I64) Concat(ys []int64) I64 {
	zs := make(I64, len(xs)+len(ys))
	copy(zs, xs)
	copy(zs[len(xs):], ys)
	return zs
}

// ConcatI64 returns a copy of the original slice with the additional boxed slice
// of values appended.
func (xs I64) ConcatI64(ys I64) I64 {
	return xs.Concat([]int64(ys))
}

// Contains checks if a value is in the list
func (xs I64) Contains(v int64) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

// Elem returns the element with the given index in the list.  Panics if the
// element does not exist.
func (xs I64) Elem(n int) int64 {
	return xs[n]
}

// Filter returns a new list of elements matching a predicate
func (xs I64) Filter(f func(int64) bool) I64 {
	ys := make(I64, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

// Head returns the first value in the list.  Panics if the list is empty.
func (xs I64) Head() int64 {
	return xs[0]
}

// Init returns a new list with all values except the last.  Panics if the
// list is empty.
func (xs I64) Init() I64 {
	ys := make(I64, len(xs)-1)
	copy(ys, xs[:len(xs)-1])
	return ys
}

// Last returns the last value in the list.  Panics if the list is empty.
func (xs I64) Last() int64 {
	return xs[len(xs)-1]
}

// Len returns the length of the list.
func (xs I64) Len() int {
	return len(xs)
}

// Less reports whether the element with index i should sort before the
// element with index j.
func (xs I64) Less(i, j int) bool {
	return xs[i] < xs[j]
}

// Map returns a new list with every element transformed by a function
func (xs I64) Map(f func(int64) int64) I64 {
	ys := make(I64, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// Reverse returns a copy of the list with the order of elements reversed.
func (xs I64) Reverse() I64 {
	ys := make(I64, len(xs))
	n := len(xs) - 1
	for i, v := range xs {
		ys[n-i] = v
	}
	return ys
}

// Swap does an in-place swap of the elements with indexes i and j.  Panics if
// the elements don't exist.  It returns nothing per the Swap signature of
// Sort.Interface.
func (xs I64) Swap(i, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}

// Tail returns a new list with all values except the head.  Panics if the
// list is empty.
func (xs I64) Tail() I64 {
	ys := make(I64, len(xs)-1)
	copy(ys, xs[1:len(xs)])
	return ys
}

// Unbox returns the list as a native slice
func (xs I64) Unbox() []int64 {
	return []int64(xs)
}
