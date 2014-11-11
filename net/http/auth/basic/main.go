// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "encoding/base64"
    "errors"
    "flag"
    "log"
    "net/http"
    "strings"
)

var (
    ErrInvalidAuth = errors.New("Invalid auth")

    login    = flag.String("login", "", "Define the login of the authentication")
    password = flag.String("password", "", "Define the password of the authentication")
    realm    = flag.String("realm", "", "Define the message of the prompt of the authentication")
    addr     = flag.String("addr", "", "addr host:port")
)

type BasicAuth struct {
    Login    string
    Password string
    Realm    string
}

func NewHttpBasicAuth(login, pass string) *BasicAuth {
    return &BasicAuth{Login: login, Password: pass}
}

func (a *BasicAuth) Authenticate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("WWW-Authenticate", "Basic realm=\""+a.Realm+"\"")
    http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func (a *BasicAuth) BasicAuthHandler(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if a.ValidAuth(r) != nil {
            a.Authenticate(w, r)
        }
    })
}

func (a *BasicAuth) ValidAuth(r *http.Request) error {
    s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
    if len(s) != 2 || s[0] != "Basic" {
        return ErrInvalidAuth
    }

    b, err := base64.StdEncoding.DecodeString(s[1])
    if err != nil {
        return err
    }

    pair := strings.SplitN(string(b), ":", 2)
    if len(pair) != 2 {
        return ErrInvalidAuth
    }

    if a.Login == pair[0] && a.Password == pair[1] {
        return nil
    }

    return ErrInvalidAuth
}

func main() {
    flag.Parse()

    auth := NewHttpBasicAuth(*login, *password)
    fs := http.FileServer(http.Dir("/"))
    handler := auth.BasicAuthHandler(fs)

    err := http.ListenAndServe(*addr, handler)
    if err != nil {
        log.Fatalf(err.Error())
    }
}
