package main

import "fmt"

func modifyValue(x *int) {
	*x = 20}

func main() {
	var val int
	val = 10

	var ptr *int
	ptr = &val

	fmt.Println("Value of val:", val, " Address of val:", ptr)
	fmt.Println("Data in pointer is", *ptr)

	var ptr2 *int
	fmt.Println("Value of ptr2:", ptr2)


	x:=10
	fmt.Println("Value of x:", x)
	modifyValue(&x)
	fmt.Println("Value of x after modification:", x)
}
