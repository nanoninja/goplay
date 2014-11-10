// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"time"
)

// openssl genrsa -out demo.rsa 1024 # the 1024 is the size of the key we are generating
// openssl rsa -in demo.rsa -pubout > demo.rsa.pub
var (
    privateKey []byte
    publicKey  []byte
)

func init() {
    privateKey, _ = ioutil.ReadFile("keys/demo.rsa")
    publicKey, _ = ioutil.ReadFile("keys/demo.rsa.pub")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
    token := jwt.New(jwt.GetSigningMethod("RS256"))
    token.Claims["ID"] = "This is my super fake ID"
    token.Claims["exp"] = time.Now().Unix() + 30
    tokenString, _ := token.SignedString(privateKey)

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `{"token": %s}`, tokenString)
}

func jwtHandler(w http.ResponseWriter, r *http.Request) {
    token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
        return publicKey, nil
    })

    fmt.Println(err)
    fmt.Println(token)

    if err != nil {
        http.Error(w, "Bad request", 500)
    }

	if token.Valid {
        fmt.Fprintf(w, "Yo man! ")
	} else {
        fmt.Fprintf(w, "Bad ! ")
	}
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	login, password := r.FormValue("login"), r.FormValue("password")

	if login == "foo" && password == "secret" {
		fmt.Fprintf(w, "Connect\n")
		// jwtHandler(w, r)
	} else {
		fmt.Fprintf(w, "Connect\n")
	}

	fmt.Fprintf(w, "Server start...\n")
	fmt.Fprintf(w, login)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/login", http.HandlerFunc(handleLogin))
	mux.Handle("/api", http.HandlerFunc(authHandler))
	http.ListenAndServe(":8080", mux)
}
