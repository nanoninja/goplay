// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "fmt"
    "io"
    "log"
    "net"
    "os"
)

type consoleWriter struct {
    c net.Conn
    w io.Writer
}

func (w consoleWriter) Write(v []byte) (int, error) {
    fmt.Fprintf(w.w, "%v > ", w.c.RemoteAddr())
    return w.w.Write(v)
}

func main() {
    l, err := net.Listen("tcp", "127.0.0.1:8080")
    if err != nil {
        log.Fatal(err)
    }
    defer l.Close()

    log.Println("Listening on", l.Addr())

    for {
        conn, err := l.Accept()
        if err != nil {
            log.Fatal(err)
            continue
        }
        go io.Copy(consoleWriter{conn, os.Stdout}, conn)
    }
}
