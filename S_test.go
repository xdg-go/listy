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
			Input []string
			Unbox []string
		}
	}
	getSTestData(is, &data)

	is.Equal(listy.S(data.Box.Input).Unbox(), data.Box.Unbox)
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
