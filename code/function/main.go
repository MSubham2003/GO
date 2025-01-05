package main

import "fmt"

func simpleFunction() {
	fmt.Println("THis is a simple function")
}

func add(a, b, c int) int {
	return a + b + c
}

func main() {
	fmt.Println("Function call and make function in go")
	simpleFunction()

	result := add(10,20,30)
	fmt.Println(result)
}
