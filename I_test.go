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

func getITestData(is *testy.T, v interface{}) {
	err := json.Unmarshal(testdata["ints.json"], v)
	if err != nil {
		is.Fatalf("Error unmarshaling ints.json: %s", err)
	}
}

func TestListIBox(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Box struct {
			Input []int
			Unbox []int
		}
	}
	getITestData(is, &data)

	is.Equal(listy.I(data.Box.Input).Unbox(), data.Box.Unbox)
}

func TestListIContains(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Contains struct {
			Input         []int
			ContainsTrue  []int
			ContainsFalse []int
		}
	}
	getITestData(is, &data)

	xs := listy.I(data.Contains.Input)

	for _, x := range data.Contains.ContainsTrue {
		is.True(xs.Contains(x))
	}
	for _, x := range data.Contains.ContainsFalse {
		is.False(xs.Contains(x))
	}
}

func TestListIFilter(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Filter struct {
			Input    []int
			Lessthan int
			Filtered []int
		}
	}
	getITestData(is, &data)

	xs := listy.I(data.Filter.Input)

	ys := xs.Filter(func(v int) bool { return v < data.Filter.Lessthan })

	is.Equal(ys.Unbox(), data.Filter.Filtered)
}

func TestListIMap(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Map struct {
			Input  []int
			Add    int
			Mapped []int
		}
	}
	getITestData(is, &data)

	xs := listy.I(data.Map.Input)
	ys := xs.Map(func(v int) int { return v + data.Map.Add })

	is.Equal(ys.Unbox(), data.Map.Mapped)
}
