package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 256
	fmt.Println(intNum)

	var floatNum float32 = 256.5
	fmt.Println(floatNum)

	fmt.Println(float32(intNum) + floatNum)

	var myString string = `Hello Bekah!`
	fmt.Println(myString)

	//get string length with utf8.RuneCountInString(myString)
	fmt.Print(utf8.RuneCountInString(myString))

	var myBoolean bool = false
	fmt.Println("")
	fmt.Println(myBoolean)

	//Get fancy and initalize multiple variables at once with infered data type.
	//(This is only safe when the data-type is obvious.)
	//Readability and accessibility are more important than fanciness, always
	var1, var2 := 3, 3.14
	//these are stored as int and float64, according to error when I tried to add them
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(float64(var1) + var2)
}
