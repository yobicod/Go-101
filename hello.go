package main

import (
	"fmt"
	"html"
)

func main() {
	fmt.Println("Hello")
	fmt.Println("Go-101")
	fmt.Print("Hello")
	fmt.Print("Go-101")

	fmt.Println(true && false)

	fmt.Println(`new
	line`)

	const htmlCode = "<p>Hello Go-101<p/>"
	en := html.EscapeString(htmlCode)
	de := html.UnescapeString(htmlCode)
	fmt.Println(en)
	fmt.Println(de)

}
