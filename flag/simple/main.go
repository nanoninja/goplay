// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// A simple flag

// Define flags using flag.String():
//  flag.String("yourFlagName", "yourDefaultValue", "Your message of help")
//
// Command line flag syntax:
//  go run main.go -username=bar
//
// With Go build:
//  go build main.go
//  ./main -username=bar

package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "Foo", "Enter your username")
	flag.Parse()
	fmt.Println("Your username is: ", *username)
}
