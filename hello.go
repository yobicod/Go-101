package main

import (
	"GO-101/math"
	"fmt"
	"html"
	"log"
	"net/url"
	"strconv"
)

func main() {

	// loop
	for i := 1; i < 10; i++ {
		if i > 5 {
			fmt.Println("Stop loop process")
			break
		}
		// run this if others succesfully execute
		defer fmt.Printf("Round %d valus is %d\n", i, i)
	}

	fmt.Println("Hello")
	fmt.Println("Go-101")
	fmt.Print("Hello")
	fmt.Print("Go-101")

	fmt.Println(true && false)

	fmt.Println(`new
	line`)

	// Encode escape char(html)
	const htmlCode = "<p>Hello Go-101<p/>"
	en := html.EscapeString(htmlCode)
	de := html.UnescapeString(htmlCode)
	fmt.Println(en)
	fmt.Println(de)

	// Encode escape char(URL)
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

	// Formatted string
	var value int16 = 250
	fmt.Printf("type of value is %T\n", value)

	fmt.Printf("Value of str is %v and type is %T\n", str, str)

	// Get input
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

	// Convert int to float
	var num1 float32 = 1.5
	var num2 int = 4
	sum := num1 + float32(num2)
	fmt.Println(sum)

	// Convert numeric string to int
	var numericString string = "18"
	convNumericString, err := strconv.Atoi(numericString)
	if err == nil {
		fmt.Printf("%d : type is %T", convNumericString, convNumericString)
	} else {
		fmt.Println(err)
	}

	var numeric int = 18

	// Convert int to string
	conStringNumeric := strconv.Itoa(numeric)
	fmt.Printf("Value is %s, Type is %T\n", conStringNumeric, conStringNumeric)

	// Array declaration with fix size
	var salary = [5]int{1, 2, 3, 4, 5}
	fmt.Println(salary)

	// assign index 1 = 4, 6 = 10
	arr := [10]int{1: 4, 6: 10}
	fmt.Println(arr)

	// Array declaration with free-size
	dynamicArray := [...]int{1, 2, 3}
	fmt.Println(dynamicArray)

	// loop through array
	acc := 0
	for i := 0; i < len(dynamicArray); i++ {
		acc += dynamicArray[i]
	}
	fmt.Println("Acc is " + strconv.Itoa(acc))

	// loop range
	for index, value := range dynamicArray {
		fmt.Println(index)
		acc += value
	}
	fmt.Println(acc)

	sumFunctionResult := calcSum(arr)
	fmt.Println("Result from called sum function is " + strconv.Itoa(sumFunctionResult))

	// Slice
	numbers := []int{1, 2, 3, 4, 5, 6}
	numbers = append(numbers, 7, 8, 9, 10)
	fmt.Println(numbers)
	slices := []int{1, 2, 3}
	fmt.Println(slices)

	// Map
	z := map[string]string{"name": "Visal", "age": "22", "gender": "male"}
	fmt.Println(z["name"])
	// delete key
	delete(z, "gender")
	fmt.Println(z)
	// check how many key
	fmt.Println(len(z))
	//check key exist in map (val if exist return value else 0, status return boolean)
	val, status := z["gender"]
	if status {
		fmt.Println(val)
	} else {
		fmt.Println("Not found")
	}

	// pass by value
	a := 200
	fmt.Println("Before call fn", a)
	changeByValue(a)
	fmt.Println("after call fn", a)
	fmt.Println("*************")
	fmt.Println("Before call fn", a)
	changeByReference(&a)
	fmt.Println("after call fn", a)

	answer := math.Add(1, 2)
	fmt.Printf("Answer is %d\n", answer)
}

func calcSum(arr [10]int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func changeByValue(number int) {
	fmt.Println("Before change fn", number)
	number = 500
	fmt.Println("after change fn", number)
}

// set pointer to address of args
func changeByReference(number *int) {
	fmt.Println("Before change fn", *number)
	*number = 500
	fmt.Println("after change fn", *number)
}
