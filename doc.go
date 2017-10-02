// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// Package listy provides functional programming idioms for Go slices.  The
// list of functions is inspired by a subset of the Data.List package from
// Haskell, and similar libraries seen in other languages.
//
// The emphasis of this library is on expressiveness and clarity.  With one
// exception (`Swap`, for `Sort.Interface`), methods treat inputs as
// immutable, returning copies rather than mutating slices in place.
package listy

//go:generate go run generator.go
