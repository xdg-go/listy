// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// +build ignore

// This file generates code for different types based on the I8.go and
// I8_test.go files.

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
)

type generated struct {
	TypeName string
	BaseType string
	TestFile string
}

// Key is listy type; value is native type
var newTypes = []generated{
	{TypeName: "I", BaseType: "int", TestFile: "ints.json"},
	{TypeName: "U", BaseType: "uint", TestFile: "ints.json"},
	{TypeName: "U8", BaseType: "uint8", TestFile: "ints.json"},
	{TypeName: "U16", BaseType: "uint16", TestFile: "ints.json"},
	{TypeName: "U32", BaseType: "uint32", TestFile: "ints.json"},
	{TypeName: "U64", BaseType: "uint64", TestFile: "ints.json"},
	{TypeName: "I16", BaseType: "int16", TestFile: "ints.json"},
	{TypeName: "I32", BaseType: "int32", TestFile: "ints.json"},
	{TypeName: "I64", BaseType: "int64", TestFile: "ints.json"},
	{TypeName: "S", BaseType: "string", TestFile: "strings.json"},
}

func (g generated) replaceTypes(orig []byte) (new []byte) {
	return bytes.Replace(
		bytes.Replace(
			bytes.Replace(
				orig,
				[]byte("I8"), []byte(g.TypeName), -1,
			),
			[]byte("int8"), []byte(g.BaseType), -1,
		),
		[]byte("ints.json"), []byte(g.TestFile), -1,
	)
}

func main() {
	log.SetFlags(0)

	// read in I8.go and I8_test.go as our "templates"
	go_file, err := ioutil.ReadFile("I8.go")
	if err != nil {
		log.Fatal("Couldn't read I8.go")
	}

	test_file, err := ioutil.ReadFile("I8_test.go")
	if err != nil {
		log.Fatal("Couldn't read I8_test.go")
	}

	// loop over the map of types; for each type, copy contents of I8/I8_test
	// and swap in new listy/native types to a new file.

	for _, v := range newTypes {
		log.Printf("Generating files for %s\n", v.TypeName)
		new_go_file_name := fmt.Sprintf("%s.go", v.TypeName)
		new_test_file_name := fmt.Sprintf("%s_test.go", v.TypeName)

		new_go_file := formatFile(new_go_file_name, v.replaceTypes(go_file))
		new_test_file := formatFile(new_test_file_name, v.replaceTypes(test_file))

		writeFile(new_go_file_name, new_go_file)
		writeFile(new_test_file_name, new_test_file)
	}

}

// format returns code cleaned up by gofmt or original if there is an error
func formatFile(name string, b []byte) []byte {
	src, err := format.Source(b)
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated in %s: %s", name, err)
		log.Printf("warning: compile the package to analyze the error")
		return b
	}
	return src
}

func writeFile(name string, data []byte) {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString("// DO NOT EDIT!  This file was generated via `go generate`\n\n")
	f.Write(data)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
