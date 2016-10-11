// Author: Gareth Lynskey
// Date: October 6th, 2016
// Adapted from: https://go-macaron.com

package main

import "gopkg.in/macaron.v1"

func main() {
    m := macaron.Classic() //import macaron package
    m.Use(macaron.Renderer()) //renders the html page from templates

    m.Get("/macaron", func() string {
        return "Hello world Macaron!"
    })

    m.Get("/reverse/:name", func(ctx *macaron.Context) {
        ctx.Data["Name"] = Reverse(ctx.Params(":name"))
        ctx.HTML(200, "hello") // 200 is the response code.
    })

    m.Run()
}

func Reverse(s string) string {
	// Here we use runes in the reverse function to reverse a string.
	//each character is taken as an individual element so we can loop through them
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}