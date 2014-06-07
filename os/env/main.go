// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Different ways to set a key/value pair on the system
// It will depend of on your machine
// FOO=20
// export FOO=20
// set FOO=20

package main

import (
	"fmt"
	"os"
)

func main() {
	// Set a key/value pair and display
	os.Setenv("BAR", "10")
	fmt.Println("The value of BAR is: ", os.Getenv("BAR"))

	// Get the value of FOO environment variable
	fmt.Println("The value of FOO is: ", os.Getenv("FOO"))
}
