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

func getSTestData(is *testy.T, v interface{}) {
	err := json.Unmarshal(testdata["strings.json"], v)
	if err != nil {
		is.Fatalf("Error unmarshaling strings.json: %s", err)
	}
}

func TestListSBox(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Box struct {
			Concat  []string
			Elem3   string
			Head    string
			Init    []string
			Input   []string
			Last    string
			Len     int
			Less23  bool
			Less32  bool
			Reverse []string
			Swap    []string
			Tail    []string
			Unbox   []string
		}
	}
	getSTestData(is, &data)

	xs := listy.S(data.Box.Input)

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
	is.Equal(xs.ConcatS(listy.S(data.Box.Input)).Unbox(), data.Box.Concat)
	is.Equal(xs.Reverse().Unbox(), data.Box.Reverse)

	// For Sort.Interface, Swap is in place and returns nothing
	xs = listy.S(data.Box.Input)
	xs.Swap(0, 1)
	is.Equal(xs.Unbox(), data.Box.Swap)
}

func TestListSContains(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Contains struct {
			Input         []string
			ContainsTrue  []string
			ContainsFalse []string
		}
	}
	getSTestData(is, &data)

	xs := listy.S(data.Contains.Input)

	for _, x := range data.Contains.ContainsTrue {
		is.True(xs.Contains(x))
	}
	for _, x := range data.Contains.ContainsFalse {
		is.False(xs.Contains(x))
	}
}

func TestListSFilter(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Filter struct {
			Input    []string
			Lessthan string
			Filtered []string
		}
	}
	getSTestData(is, &data)

	xs := listy.S(data.Filter.Input)

	ys := xs.Filter(func(v string) bool { return v < data.Filter.Lessthan })

	is.Equal(ys.Unbox(), data.Filter.Filtered)
}

func TestListSMap(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Map struct {
			Input  []string
			Add    string
			Mapped []string
		}
	}
	getSTestData(is, &data)

	xs := listy.S(data.Map.Input)
	ys := xs.Map(func(v string) string { return v + data.Map.Add })

	is.Equal(ys.Unbox(), data.Map.Mapped)
}

func TestListSUniq(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Uniq struct {
			Cases map[string][2][]string
		}
	}
	getSTestData(is, &data)

	for k, v := range data.Uniq.Cases {
		is.Label(k).Equal(listy.S(v[0]).Uniq().Unbox(), v[1])
	}

}
