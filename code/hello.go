package main

import (
	"fmt"
	"learn/subpack"
)

func main() {
	fmt.Println("Hello")
	subpack.PrintMessage("hello world")

	var f_name string = "Subham"
	var l_name = "Mohanty"
	fmt.Println(f_name)
	fmt.Println(l_name)

	var x int = 10
	var y = 20
	fmt.Println("x:", x, ",y:", y)

	var amt float64 = 2150.35
	var rbt = 10.25
	fmt.Println("amount:", amt, "\nRebet:", rbt)

	var decided bool = true
	var value = false
	fmt.Println(decided,value)

	const pi = 3.14
	fmt.Println(pi)

	a := 100
	b := "Hello"
	fmt.Println(a,b)


	// Private variable
	var name = "abc"

	// if 1st letter is capital then it is public means it can be accessed outside the package
	// thats why we write Println instead of println
	
	// Public variable   
	var Name = "abc"
	fmt.Println(name,Name)
}
