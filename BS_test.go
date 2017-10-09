// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/xdg/listy"
	"github.com/xdg/testy"
)

func getBSTestData(is *testy.T, v interface{}) {
	err := json.Unmarshal(testdata["base64.json"], v)
	if err != nil {
		is.Fatalf("Error unmarshaling base64.json: %s", err)
	}
}

func TestListBSBox(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Box struct {
			Concat  [][]byte
			Elem3   []byte
			Head    []byte
			Init    [][]byte
			Input   [][]byte
			Last    []byte
			Len     int
			Less23  bool
			Less32  bool
			Reverse [][]byte
			Swap    [][]byte
			Tail    [][]byte
			Unbox   [][]byte
		}
	}
	getBSTestData(is, &data)

	xs := listy.BS(data.Box.Input)

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
	is.Equal(xs.ConcatBS(listy.BS(data.Box.Input)).Unbox(), data.Box.Concat)
	is.Equal(xs.Reverse().Unbox(), data.Box.Reverse)

	// For Sort.Interface, Swap is in place and returns nothing
	xs = listy.BS(data.Box.Input)
	xs.Swap(0, 1)
	is.Equal(xs.Unbox(), data.Box.Swap)
}

func TestListBSContains(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Contains struct {
			Input         [][]byte
			ContainsTrue  [][]byte
			ContainsFalse [][]byte
		}
	}
	getBSTestData(is, &data)

	xs := listy.BS(data.Contains.Input)

	for _, x := range data.Contains.ContainsTrue {
		is.True(xs.Contains(x))
	}
	for _, x := range data.Contains.ContainsFalse {
		is.False(xs.Contains(x))
	}
}

func TestListBSFilter(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Filter struct {
			Input    [][]byte
			Lessthan []byte
			Filtered [][]byte
		}
	}
	getBSTestData(is, &data)

	xs := listy.BS(data.Filter.Input)

	ys := xs.Filter(func(v []byte) bool { return bytes.Compare(v, data.Filter.Lessthan) == -1 })

	is.Equal(ys.Unbox(), data.Filter.Filtered)
}

func TestListBSMap(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Map struct {
			Input  [][]byte
			Add    []byte
			Mapped [][]byte
		}
	}
	getBSTestData(is, &data)

	xs := listy.BS(data.Map.Input)
	ys := xs.Map(func(v []byte) []byte { return append(v, data.Map.Add...) })

	is.Equal(ys.Unbox(), data.Map.Mapped)
}

func TestListBSUniq(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Uniq struct {
			Cases map[string][2][][]byte
		}
	}
	getBSTestData(is, &data)

	for k, v := range data.Uniq.Cases {
		is.Label(k).Equal(listy.BS(v[0]).Uniq().Unbox(), v[1])
	}

}
