// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var testdata map[string][]byte
var testfiles = []string{"ints.json", "strings.json"}

func TestMain(m *testing.M) {
	testdata = make(map[string][]byte)
	for _, s := range testfiles {
		b, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s", s))
		if err != nil {
			log.Fatalf("Error reading %s: %s", s, err)
		}
		testdata[s] = b
	}

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
