// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"net/http"
	"strings"
)

var (
	login    string
	password string
	realm    string
	port     int
)

func init() {
	flag.StringVar(&login, "login", "", "Define the login of the authentication")
	flag.StringVar(&password, "password", "", "Define the password of the authentication")
	flag.StringVar(&realm, "realm", "Authenticate", "Define the message of the prompt of the authentication")
	flag.IntVar(&port, "port", 8080, "Define the port of the server")
	flag.Parse()
}

type BasicAuthHttp struct {
	Login    string
	Password string
	Realm    string
	Handler  http.Handler
}

func NewHttpBasicAuth(login, pass string) *BasicAuthHttp {
	return &BasicAuthHttp{Login: login, Password: pass}
}

func (a *BasicAuthHttp) Authenticate(w http.ResponseWriter, code int) {
	w.Header().Set("WWW-Authenticate", "Basic realm=\""+a.Realm+"\"")
	http.Error(w, http.StatusText(code), code)
}

func (a *BasicAuthHttp) BasicAuthHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header["Authorization"]; len(auth) > 0 {
			token := strings.Replace(string(auth[0]), "Basic ", "", 1)
			s, err := base64.StdEncoding.DecodeString(token)
			if err != nil {
				fmt.Println(err)
			}
			if parts := strings.Split(string(s), ":"); len(parts) == 2 {
				if a.Login == parts[0] && a.Password == parts[1] {
					h.ServeHTTP(w, r)
				} else {
					a.Authenticate(w, 401)
				}
			}
		} else {
			a.Authenticate(w, 401)
		}
	})
}

func main() {
	auth := NewHttpBasicAuth(login, password)
	fs := http.FileServer(http.Dir("/"))
	handler := auth.BasicAuthHandler(fs)
	http.ListenAndServe(fmt.Sprintf("%s:%d", "", port), handler)
}
