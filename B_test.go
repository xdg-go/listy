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

func getBTestData(is *testy.T, v interface{}) {
	err := json.Unmarshal(testdata["ints.json"], v)
	if err != nil {
		is.Fatalf("Error unmarshaling ints.json: %s", err)
	}
}

func TestListBBox(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Box struct {
			Input  []byte
			Head   byte
			Elem3  byte
			Tail   []byte
			Init   []byte
			Last   byte
			Len    int
			Less23 bool
			Less32 bool
			Swap   []byte
			Unbox  []byte
		}
	}
	getBTestData(is, &data)

	xs := listy.B(data.Box.Input)

	is.Equal(xs.Unbox(), data.Box.Unbox)
	is.Equal(xs.Head(), data.Box.Head)
	is.Equal(xs.Elem(3), data.Box.Elem3)
	is.Equal(xs.Tail().Unbox(), data.Box.Tail)
	is.Equal(xs.Init().Unbox(), data.Box.Init)
	is.Equal(xs.Last(), data.Box.Last)
	is.Equal(xs.Len(), data.Box.Len)
	is.Equal(xs.Less(2, 3), data.Box.Less23)
	is.Equal(xs.Less(3, 2), data.Box.Less32)

	// For Sort.Interface, Swap returns nothing
	xs.Swap(0, 1)
	is.Equal(xs.Unbox(), data.Box.Swap)
}

func TestListBContains(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Contains struct {
			Input         []byte
			ContainsTrue  []byte
			ContainsFalse []byte
		}
	}
	getBTestData(is, &data)

	xs := listy.B(data.Contains.Input)

	for _, x := range data.Contains.ContainsTrue {
		is.True(xs.Contains(x))
	}
	for _, x := range data.Contains.ContainsFalse {
		is.False(xs.Contains(x))
	}
}

func TestListBFilter(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Filter struct {
			Input    []byte
			Lessthan byte
			Filtered []byte
		}
	}
	getBTestData(is, &data)

	xs := listy.B(data.Filter.Input)

	ys := xs.Filter(func(v byte) bool { return v < data.Filter.Lessthan })

	is.Equal(ys.Unbox(), data.Filter.Filtered)
}

func TestListBMap(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Map struct {
			Input  []byte
			Add    byte
			Mapped []byte
		}
	}
	getBTestData(is, &data)

	xs := listy.B(data.Map.Input)
	ys := xs.Map(func(v byte) byte { return v + data.Map.Add })

	is.Equal(ys.Unbox(), data.Map.Mapped)
}
