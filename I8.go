// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy

// I8 wraps a slice of int8
type I8 []int8

type int8ToBool func(int8) bool

type int8ToInt func(int8) int8

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
func (xs I8) Filter(f int8ToBool) I8 {
	ys := make(I8, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

// Map returns a new list with every element transformed by a function
func (xs I8) Map(f int8ToInt) I8 {
	ys := make(I8, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// Unbox returns the list as a native slice
func (xs I8) Unbox() []int8 {
	return []int8(xs)
}
