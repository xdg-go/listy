// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy

import "bytes"

// BS wraps a slice of []byte
type BS [][]byte

// Concat returns a copy of the original slice with the additional slice of
// values appended.
func (xs BS) Concat(ys [][]byte) BS {
	zs := make(BS, len(xs)+len(ys))
	copy(zs, xs)
	copy(zs[len(xs):], ys)
	return zs
}

// ConcatBS returns a copy of the original slice with the additional boxed slice
// of values appended.
func (xs BS) ConcatBS(ys BS) BS {
	return xs.Concat([][]byte(ys))
}

// Contains checks if a value is in the list
func (xs BS) Contains(v []byte) bool {
	for _, x := range xs {
		if bytes.Equal(x, v) {
			return true
		}
	}
	return false
}

// Elem returns the element with the given index in the list.  Panics if the
// element does not exist.
func (xs BS) Elem(n int) []byte {
	return xs[n]
}

// Filter returns a new list of elements matching a predicate
func (xs BS) Filter(f func([]byte) bool) BS {
	ys := make(BS, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

// Head returns the first value in the list.  Panics if the list is empty.
func (xs BS) Head() []byte {
	return xs[0]
}

// Init returns a new list with all values except the last.  Panics if the
// list is empty.
func (xs BS) Init() BS {
	ys := make(BS, len(xs)-1)
	copy(ys, xs[:len(xs)-1])
	return ys
}

// Last returns the last value in the list.  Panics if the list is empty.
func (xs BS) Last() []byte {
	return xs[len(xs)-1]
}

// Len returns the length of the list.
func (xs BS) Len() int {
	return len(xs)
}

// Less reports whether the element with index i should sort before the
// element with index j.
func (xs BS) Less(i, j int) bool {
	return bytes.Compare(xs[i], xs[j]) == -1
}

// Map returns a new list with every element transformed by a function
func (xs BS) Map(f func([]byte) []byte) BS {
	ys := make(BS, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// Reverse returns a copy of the list with the order of elements reversed.
func (xs BS) Reverse() BS {
	ys := make(BS, len(xs))
	n := len(xs) - 1
	for i, v := range xs {
		ys[n-i] = v
	}
	return ys
}

// Swap does an in-place swap of the elements with indexes i and j.  Panics if
// the elements don't exist.  It returns nothing per the Swap signature of
// Sort.Interface.
func (xs BS) Swap(i, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}

// Tail returns a new list with all values except the head.  Panics if the
// list is empty.
func (xs BS) Tail() BS {
	ys := make(BS, len(xs)-1)
	copy(ys, xs[1:len(xs)])
	return ys
}

// Unbox returns the list as a native slice
func (xs BS) Unbox() [][]byte {
	return [][]byte(xs)
}

// Uniq returns a new list with duplicate elements removed.
func (xs BS) Uniq() BS {
	ys := make(BS, 0)
	seen := make(map[string]struct{})
	for _, x := range xs {
		if _, ok := seen[string(x)]; ok {
			continue
		}
		ys = append(ys, x)
		seen[string(x)] = struct{}{}
	}
	return ys
}
