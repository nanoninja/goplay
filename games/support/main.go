// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is a games support
// Ask a question and get an answer support
// Just type "bye" to quit

package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "time"
)

func Welcome() {
    fmt.Println("Welcome to the games support.")
    fmt.Println("")
    fmt.Println("Write your question.")
    fmt.Println("We will try to solve your problem.")
    fmt.Println("Type 'bye' for quit the support.")
}

func Goodbye() {
    fmt.Println("We were happy to help. Bye Bye...")
}

type Responses []string

type Support struct {
    Name string
    Resp Responses
}

func NewSupport(name string) *Support {
    return &Support{name, make(Responses, 0)}
}

func (s *Support) AddResponse(str string) *Support {
    s.Resp = append(s.Resp, str)
    return s
}

func (s Support) GetResponses() Responses {
    return s.Resp
}

func (s Support) GenerateResp() string {
    rand.Seed(time.Now().Unix())
    responses := s.Resp
    return responses[rand.Intn(len(responses))]
}

func (s *Support) Init() {
    s.AddResponse("It seems strange. Could you describe the problem more precisely?")
    s.AddResponse("No other customer has ever complained about this.")
    s.AddResponse("What is your system configuration?")
    s.AddResponse("That sounds interesting. Tell me more ...")
    s.AddResponse("This is explained in the manual")
    s.AddResponse("Your description is not clear. Will he an expert you could describe this more precisely?")
    s.AddResponse("This is not a bug but a feature!")
    s.AddResponse("Could you clarify?")
}

func (s *Support) play() {
    s.Init()

    Welcome()
    var input string
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        input = scanner.Text()
        if input == "bye" {
            Goodbye()
            break
        } else {
            fmt.Println(s.GenerateResp())
        }
    }
}

func main() {
    support := NewSupport("Games Support")
    support.play()
}
