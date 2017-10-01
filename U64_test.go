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
	err := json.Unmarshal(testdata["ints"], v)
	if err != nil {
		is.Fatalf("Error unmarshaling ints.json: %s", err)
	}
}

func TestListU64Box(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Box struct {
			Input []uint64
			Unbox []uint64
		}
	}
	getU64TestData(is, &data)

	is.Equal(listy.U64(data.Box.Input).Unbox(), data.Box.Unbox)
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
