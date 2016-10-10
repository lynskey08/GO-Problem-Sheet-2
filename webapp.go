package main

import "gopkg.in/macaron.v1"

func main() {
    m := macaron.Classic()
    m.Get("/macaron", func() string {
        return "Hello world Macaron!"
    })

    m.Get("/reverse", func() string {
        return "Hello world Reverse!"
    })
    m.Run()
}