// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "bufio"
    "log"
    "net"
    "os"
)

func main() {
    addr := "127.0.0.1:8080"

    conn, err := net.Dial("tcp", addr)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Dial with %s\n", addr)
    dial(conn)
}

func dial(conn net.Conn) {
    defer conn.Close()
    scan := bufio.NewScanner(os.Stdin)

    for scan.Scan() {
        input := scan.Text() + "\n"
        conn.Write([]byte(input))
    }
}
