package main

import (
	"fmt"
	"html"
	"log"
	"net/url"
	"strconv"
)

func main() {

	for i := 1; i < 10; i++ {
		if i > 5 {
			fmt.Println("Stop loop process")
			break
		}
		defer fmt.Printf("Round %d valus is %d\n", i, i)
	}

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

	// var firstName string
	// var myAge int
	// fmt.Println("Enter your name and age : ")
	// n, err := fmt.Scan(&firstName, &myAge)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("Your first name is %v\n", firstName)
	// fmt.Printf("Age is %d", myAge)
	// fmt.Println(n)

	var num1 float32 = 1.5
	var num2 int = 4
	sum := num1 + float32(num2)
	fmt.Println(sum)

	var numericString string = "18"
	convNumericString, err := strconv.Atoi(numericString)
	if err == nil {
		fmt.Printf("%d : type is %T", convNumericString, convNumericString)
	} else {
		fmt.Println(err)
	}

	var numeric int = 18
	conStringNumeric := strconv.Itoa(numeric)
	fmt.Printf("Value is %s, Type is %T\n", conStringNumeric, conStringNumeric)

}
