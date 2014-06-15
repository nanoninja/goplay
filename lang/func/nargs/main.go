// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Give variable size arguments to a function
// http://play.golang.org/p/hbU5_B-Zr7
//
// For example:
// func Name(a ...int)
// func Name(a ...string)
// func Name(a ...interface{})

package main

import "fmt"

// A simple example with the Sum function
func Sum(a ...float64) float64 {
    var res float64

    for _, v := range a {
        res += v
    }
    return res
}

func main() {
    fmt.Println("The sum is: ", Sum(10.90, 59.50, 30, 90.99)) // output = 191.39
}
