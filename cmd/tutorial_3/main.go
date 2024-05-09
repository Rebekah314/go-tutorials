package main

import (
	"fmt"
	"strconv"
)

//import "strconv" and use strconv.Itoa(myInt) to convert myInt to a string

func main() {

	primeCount := 1
	//maxInt := 1000
	var maxInt int

	fmt.Println("Let's find all prime numbers less than a max value of your choice!")
	fmt.Println("max value: ")
	fmt.Scanln(&maxInt)

	fmt.Println("*********")
	fmt.Println("Now printing all prime numbers between 2 and " + strconv.Itoa(maxInt))
	fmt.Println(2)

	for index := 2; index < maxInt; index += 1 {

		for divisor := 2; divisor < index; divisor += 1 {
			if index%divisor == 0 {
				break
			}
			if index == divisor+1 {
				fmt.Println(index)
				primeCount += 1

			}
		}
	}
	fmt.Println("There are " + strconv.Itoa(primeCount) + " prime numbers between 2 and " + strconv.Itoa(maxInt))

}
