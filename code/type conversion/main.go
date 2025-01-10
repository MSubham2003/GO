package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Integer to float & viceversa
	var a int = 10
	fmt.Printf("Value of a: %d, Type of %d is %T\n", a, a, a)

	var b float64 = float64(a)
	fmt.Printf("Value of b: %f, Type of %f is %T\n", b, b, b)

	var x int = int(b)
	fmt.Printf("Value of x: %d, Type of %d is %T\n", x, x, x)

	//Integre to String & viceversa
	c := strconv.Itoa(a)
	fmt.Printf("Value of c: %s, Type of %s is %T\n", c, c, c)

	d, _ := strconv.Atoi(c)
	fmt.Printf("Value of d: %d, Type of %d is %T\n", d, d, d)

	//String to float & viceversa
	e := "3.14"
	f, _ := strconv.ParseFloat(e, 64)
	fmt.Printf("Value of f: %f, Type of %f is %T\n", f, f, f)

	g := strconv.FormatFloat(f, 'f', 1, 64)
	fmt.Printf("Value of g: %s, Type of %s is %T\n", g, g, g)
}
