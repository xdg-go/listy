// DO NOT EDIT!  This file was generated via `go generate`

// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy_test

import (
	"encoding/json"
	"testing"

	"github.com/xdg/listy"
	"github.com/xdg/testy"
)

func getU64TestData(is *testy.T, v interface{}) {
	err := json.Unmarshal(testdata["ints.json"], v)
	if err != nil {
		is.Fatalf("Error unmarshaling ints.json: %s", err)
	}
}

func TestListU64Box(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Box struct {
			Concat  []uint64
			Elem3   uint64
			Head    uint64
			Init    []uint64
			Input   []uint64
			Last    uint64
			Len     int
			Less23  bool
			Less32  bool
			Reverse []uint64
			Swap    []uint64
			Tail    []uint64
			Unbox   []uint64
		}
	}
	getU64TestData(is, &data)

	xs := listy.U64(data.Box.Input)

	// Tests that return ints, bools, etc. irrespective of value type
	is.Equal(xs.Len(), data.Box.Len)
	is.Equal(xs.Less(2, 3), data.Box.Less23)
	is.Equal(xs.Less(3, 2), data.Box.Less32)

	// Tests that return the value type
	is.Equal(xs.Head(), data.Box.Head)
	is.Equal(xs.Elem(3), data.Box.Elem3)
	is.Equal(xs.Last(), data.Box.Last)

	// Tests that return a slice of the value type
	is.Equal(xs.Unbox(), data.Box.Unbox)
	is.Equal(xs.Tail().Unbox(), data.Box.Tail)
	is.Equal(xs.Init().Unbox(), data.Box.Init)
	is.Equal(xs.Concat(data.Box.Input).Unbox(), data.Box.Concat)
	is.Equal(xs.ConcatU64(listy.U64(data.Box.Input)).Unbox(), data.Box.Concat)
	is.Equal(xs.Reverse().Unbox(), data.Box.Reverse)

	// For Sort.Interface, Swap is in place and returns nothing
	xs = listy.U64(data.Box.Input)
	xs.Swap(0, 1)
	is.Equal(xs.Unbox(), data.Box.Swap)
}

func TestListU64Contains(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Contains struct {
			Input         []uint64
			ContainsTrue  []uint64
			ContainsFalse []uint64
		}
	}
	getU64TestData(is, &data)

	xs := listy.U64(data.Contains.Input)

	for _, x := range data.Contains.ContainsTrue {
		is.True(xs.Contains(x))
	}
	for _, x := range data.Contains.ContainsFalse {
		is.False(xs.Contains(x))
	}
}

func TestListU64Filter(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Filter struct {
			Input    []uint64
			Lessthan uint64
			Filtered []uint64
		}
	}
	getU64TestData(is, &data)

	xs := listy.U64(data.Filter.Input)

	ys := xs.Filter(func(v uint64) bool { return v < data.Filter.Lessthan })

	is.Equal(ys.Unbox(), data.Filter.Filtered)
}

func TestListU64Map(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Map struct {
			Input  []uint64
			Add    uint64
			Mapped []uint64
		}
	}
	getU64TestData(is, &data)

	xs := listy.U64(data.Map.Input)
	ys := xs.Map(func(v uint64) uint64 { return v + data.Map.Add })

	is.Equal(ys.Unbox(), data.Map.Mapped)
}
