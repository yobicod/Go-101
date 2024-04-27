package main

import (
	"fmt"
	"html"
	"log"
	"net/url"
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

	encodePath := url.PathEscape("get/me")
	decodePath, err := url.PathUnescape(encodePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(encodePath)
	fmt.Println(decodePath)

	var str string
	var num, count int

	str = "Go language"
	num = 1
	count = 2
	fmt.Println(str, num, count)

	born := 2002
	isMale, edu := true, "Banchelor"
	fmt.Println(born, isMale, edu)

	const lightSpeed = 299792
	const Pi = 3.1451928
	distance := 5000000
	radius := 5.0
	fmt.Println("time light speed : ", distance/lightSpeed, " seconds")
	fmt.Println("Aread of cycle : ", Pi*radius*radius)

	var value int16 = 250
	fmt.Printf("type of value is %T\n", value)

	fmt.Printf("Value of str is %v and type is %T\n", str, str)
}
