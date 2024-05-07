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

	//get fancy and initalize multiple variables at once with infered data type
	var1, var2 := 3, 3.14
	fmt.Println(var1)
	fmt.Println(var2)
}
