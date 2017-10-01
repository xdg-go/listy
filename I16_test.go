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

func getI16TestData(is *testy.T, v interface{}) {
	err := json.Unmarshal(testdata["ints"], v)
	if err != nil {
		is.Fatalf("Error unmarshaling ints.json: %s", err)
	}
}

func TestListI16Box(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Box struct {
			Input []int16
			Unbox []int16
		}
	}
	getI16TestData(is, &data)

	is.Equal(listy.I16(data.Box.Input).Unbox(), data.Box.Unbox)
}

func TestListI16Contains(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Contains struct {
			Input         []int16
			ContainsTrue  []int16
			ContainsFalse []int16
		}
	}
	getI16TestData(is, &data)

	xs := listy.I16(data.Contains.Input)

	for _, x := range data.Contains.ContainsTrue {
		is.True(xs.Contains(x))
	}
	for _, x := range data.Contains.ContainsFalse {
		is.False(xs.Contains(x))
	}
}

func TestListI16Filter(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Filter struct {
			Input    []int16
			Lessthan int16
			Filtered []int16
		}
	}
	getI16TestData(is, &data)

	xs := listy.I16(data.Filter.Input)

	ys := xs.Filter(func(v int16) bool { return v < data.Filter.Lessthan })

	is.Equal(ys.Unbox(), data.Filter.Filtered)
}

func TestListI16Map(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Map struct {
			Input  []int16
			Add    int16
			Mapped []int16
		}
	}
	getI16TestData(is, &data)

	xs := listy.I16(data.Map.Input)
	ys := xs.Map(func(v int16) int16 { return v + data.Map.Add })

	is.Equal(ys.Unbox(), data.Map.Mapped)
}
