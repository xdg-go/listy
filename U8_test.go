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

func getU8TestData(is *testy.T, v interface{}) {
	err := json.Unmarshal(testdata["ints.json"], v)
	if err != nil {
		is.Fatalf("Error unmarshaling ints.json: %s", err)
	}
}

func TestListU8Box(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Box struct {
			Concat  []uint8
			Elem3   uint8
			Head    uint8
			Init    []uint8
			Input   []uint8
			Last    uint8
			Len     int
			Less23  bool
			Less32  bool
			Reverse []uint8
			Swap    []uint8
			Tail    []uint8
			Unbox   []uint8
		}
	}
	getU8TestData(is, &data)

	xs := listy.U8(data.Box.Input)

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
	is.Equal(xs.ConcatU8(listy.U8(data.Box.Input)).Unbox(), data.Box.Concat)
	is.Equal(xs.Reverse().Unbox(), data.Box.Reverse)

	// For Sort.Interface, Swap is in place and returns nothing
	xs = listy.U8(data.Box.Input)
	xs.Swap(0, 1)
	is.Equal(xs.Unbox(), data.Box.Swap)
}

func TestListU8Contains(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Contains struct {
			Input         []uint8
			ContainsTrue  []uint8
			ContainsFalse []uint8
		}
	}
	getU8TestData(is, &data)

	xs := listy.U8(data.Contains.Input)

	for _, x := range data.Contains.ContainsTrue {
		is.True(xs.Contains(x))
	}
	for _, x := range data.Contains.ContainsFalse {
		is.False(xs.Contains(x))
	}
}

func TestListU8Filter(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Filter struct {
			Input    []uint8
			Lessthan uint8
			Filtered []uint8
		}
	}
	getU8TestData(is, &data)

	xs := listy.U8(data.Filter.Input)

	ys := xs.Filter(func(v uint8) bool { return v < data.Filter.Lessthan })

	is.Equal(ys.Unbox(), data.Filter.Filtered)
}

func TestListU8Map(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Map struct {
			Input  []uint8
			Add    uint8
			Mapped []uint8
		}
	}
	getU8TestData(is, &data)

	xs := listy.U8(data.Map.Input)
	ys := xs.Map(func(v uint8) uint8 { return v + data.Map.Add })

	is.Equal(ys.Unbox(), data.Map.Mapped)
}

func TestListU8Uniq(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Uniq struct {
			Cases map[string][2][]uint8
		}
	}
	getU8TestData(is, &data)

	for k, v := range data.Uniq.Cases {
		is.Label(k).Equal(listy.U8(v[0]).Uniq().Unbox(), v[1])
	}

}
