// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "flag"
    "fmt"
    "net/http"
)

var hostname string
var port int
var root string

func init() {
    flag.StringVar(&root, "root", "/", "Define the root path. Example: /your/root/path/")
    flag.StringVar(&hostname, "hostname", "localhost", "Define the server name. Example: localhost")
    flag.IntVar(&port, "port", 8080, "Define the port of the server. Example: 80")
    flag.Parse()
}

func main() {
    serverName := fmt.Sprintf("%s:%d", hostname, port)
    err := http.ListenAndServe(serverName, http.FileServer(http.Dir(root)))
    if err != nil {
        panic(err)
    }
}
