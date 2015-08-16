// Copyright 2015 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// curl --data "username=john&password=1234" http://localhost:3000/login
// curl --request POST 'http://localhost:3000/login' --data 'username=john&password=1234'
// curl -H "Authorization: Bearer {token}" localhost:3000/hello

package main

import (
    "encoding/json"
    "jwt/auth"
    "log"
    "net/http"
)

type AuthMiddleware struct {
    next http.Handler
}

func NewAutMiddleware(handler http.Handler) *AuthMiddleware {
    return &AuthMiddleware{
        next: handler,
    }
}

func (h AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ok, err := a.IsValidToken(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
    } else {
        h.next.ServeHTTP(w, r)
    }
}

var a *auth.JWTAuth = auth.GetJWTAuthInstance()

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    requestUser := &auth.User{
        Username: r.FormValue("username"),
        Password: r.FormValue("password"),
    }
    testUser := &auth.User{
        Username: "john",
        Password: "1234",
    }

    ok, err := a.Authenticate(requestUser)(testUser)
    if err != nil {
        panic(err)
    }

    if ok {
        token, _ := a.GenerateToken("abcd")
        data, err := json.Marshal(map[string]string{
            "token": token,
        })
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        w.Write(data)
    }
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(" Go Hello "))
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(" Go Test "))
}

func main() {
    m := http.NewServeMux()

    m.HandleFunc("/login", LoginHandler)
    m.Handle("/hello", NewAutMiddleware(http.HandlerFunc(HelloHandler)))
    m.Handle("/test", NewAutMiddleware(http.HandlerFunc(TestHandler)))

    log.Fatal(http.ListenAndServe(":3000", m))
}
