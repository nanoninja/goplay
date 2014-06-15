// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// A simple example with the init function
// http://play.golang.org/p/uJNXRbdOCj

package main

import "fmt"

func init() {
	fmt.Println("1: Init function")
}

func init() {
	fmt.Println("2: Another init function")
}

func main() {
	fmt.Println("3: Main function")
}
