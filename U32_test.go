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

func getU32TestData(is *testy.T, v interface{}) {
	err := json.Unmarshal(testdata["ints.json"], v)
	if err != nil {
		is.Fatalf("Error unmarshaling ints.json: %s", err)
	}
}

func TestListU32Box(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Box struct {
			Input []uint32
			Unbox []uint32
		}
	}
	getU32TestData(is, &data)

	is.Equal(listy.U32(data.Box.Input).Unbox(), data.Box.Unbox)
}

func TestListU32Contains(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Contains struct {
			Input         []uint32
			ContainsTrue  []uint32
			ContainsFalse []uint32
		}
	}
	getU32TestData(is, &data)

	xs := listy.U32(data.Contains.Input)

	for _, x := range data.Contains.ContainsTrue {
		is.True(xs.Contains(x))
	}
	for _, x := range data.Contains.ContainsFalse {
		is.False(xs.Contains(x))
	}
}

func TestListU32Filter(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Filter struct {
			Input    []uint32
			Lessthan uint32
			Filtered []uint32
		}
	}
	getU32TestData(is, &data)

	xs := listy.U32(data.Filter.Input)

	ys := xs.Filter(func(v uint32) bool { return v < data.Filter.Lessthan })

	is.Equal(ys.Unbox(), data.Filter.Filtered)
}

func TestListU32Map(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Map struct {
			Input  []uint32
			Add    uint32
			Mapped []uint32
		}
	}
	getU32TestData(is, &data)

	xs := listy.U32(data.Map.Input)
	ys := xs.Map(func(v uint32) uint32 { return v + data.Map.Add })

	is.Equal(ys.Unbox(), data.Map.Mapped)
}
