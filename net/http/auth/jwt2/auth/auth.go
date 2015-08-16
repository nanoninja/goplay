// Copyright 2015 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://tools.ietf.org/html/rfc7519

package auth

import (
    "bufio"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    jwt "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
    "net/http"
    "os"
    "time"
)

type User struct {
    Username string
    Password string
}

type JWTAuth struct {
    PublicKey  *rsa.PublicKey
    privateKey *rsa.PrivateKey
}

var jwtAuthInstance *JWTAuth
var TokenDuration int = 20

func GetJWTAuthInstance() *JWTAuth {
    if jwtAuthInstance == nil {
        jwtAuthInstance = &JWTAuth{
            PublicKey:  getPublicKey("keys/auth.rsa.pub"),
            privateKey: getPrivateKey("keys/auth.rsa"),
        }
    }
    return jwtAuthInstance
}

func (a *JWTAuth) GenerateToken(userID string) (string, error) {
    token := jwt.New(jwt.SigningMethodRS512)
    token.Claims["exp"] = time.Now().Add(time.Hour * time.Duration(TokenDuration)).Unix()
    token.Claims["iat"] = time.Now().Unix()
    token.Claims["sub"] = userID
    tokenString, err := token.SignedString(a.privateKey)
    if err != nil {
        panic(err)
        return "", err
    }
    return tokenString, nil
}

func (a *JWTAuth) Authenticate(req *User) func(*User) (bool, error) {
    return func(user *User) (bool, error) {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
        if err != nil {
            return false, err
        }
        return req.Username == user.Username &&
            bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)) == nil, nil
    }
}

func (a *JWTAuth) IsValidToken(r *http.Request) (bool, error) {
    token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
        return a.PublicKey, nil
    })
    return token.Valid, err
}

func getPublicKey(filename string) *rsa.PublicKey {
    publicKeyFile, err := os.Open(filename)
    if err != nil {
        panic(err)
    }

    pemInfoFile, _ := publicKeyFile.Stat()
    pemBytes := make([]byte, pemInfoFile.Size())

    buf := bufio.NewReader(publicKeyFile)
    _, err = buf.Read(pemBytes)
    data, _ := pem.Decode([]byte(pemBytes))

    publicKeyFile.Close()

    publicKey, err := x509.ParsePKIXPublicKey(data.Bytes)
    if err != nil {
        panic(err)
    }

    rsaPub, ok := publicKey.(*rsa.PublicKey)
    if !ok {
        panic(err)
    }
    return rsaPub
}

func getPrivateKey(filename string) *rsa.PrivateKey {
    privateKeyFile, err := os.Open(filename)
    if err != nil {
        panic(err)
    }

    pemFileInfo, _ := privateKeyFile.Stat()
    pemBytes := make([]byte, pemFileInfo.Size())

    buf := bufio.NewReader(privateKeyFile)
    _, err = buf.Read(pemBytes)
    data, _ := pem.Decode([]byte(pemBytes))

    privateKeyFile.Close()

    privateKey, err := x509.ParsePKCS1PrivateKey(data.Bytes)
    if err != nil {
        panic(err)
    }
    return privateKey
}
