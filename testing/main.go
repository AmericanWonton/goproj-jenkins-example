package main

import "fmt"

func main() {
	fmt.Printf("This is a test print for golang.\n ")
	result := Calculate(2)
	fmt.Printf("Here is the result calculate: %v\n", result)

	result2 := Add(2, 4)
	fmt.Printf("Here is the result2 add: %v\n", result2)
}

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func Add(x, y int) int {
	return x + y
}
