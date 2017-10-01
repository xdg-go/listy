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

func getRTestData(is *testy.T, v interface{}) {
	err := json.Unmarshal(testdata["ints.json"], v)
	if err != nil {
		is.Fatalf("Error unmarshaling ints.json: %s", err)
	}
}

func TestListRBox(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Box struct {
			Input []rune
			Unbox []rune
		}
	}
	getRTestData(is, &data)

	is.Equal(listy.R(data.Box.Input).Unbox(), data.Box.Unbox)
}

func TestListRContains(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Contains struct {
			Input         []rune
			ContainsTrue  []rune
			ContainsFalse []rune
		}
	}
	getRTestData(is, &data)

	xs := listy.R(data.Contains.Input)

	for _, x := range data.Contains.ContainsTrue {
		is.True(xs.Contains(x))
	}
	for _, x := range data.Contains.ContainsFalse {
		is.False(xs.Contains(x))
	}
}

func TestListRFilter(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Filter struct {
			Input    []rune
			Lessthan rune
			Filtered []rune
		}
	}
	getRTestData(is, &data)

	xs := listy.R(data.Filter.Input)

	ys := xs.Filter(func(v rune) bool { return v < data.Filter.Lessthan })

	is.Equal(ys.Unbox(), data.Filter.Filtered)
}

func TestListRMap(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Map struct {
			Input  []rune
			Add    rune
			Mapped []rune
		}
	}
	getRTestData(is, &data)

	xs := listy.R(data.Map.Input)
	ys := xs.Map(func(v rune) rune { return v + data.Map.Add })

	is.Equal(ys.Unbox(), data.Map.Mapped)
}
