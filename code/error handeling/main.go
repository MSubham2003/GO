package main

import "fmt"

func div(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "Denominator must not be 0"
	}
	return a / b, ""
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator must not be 0")
	}
	return a / b, nil
}

func main() {
	fmt.Println("This is for Error handeling in function")

	// _ this variable is use to discard value we get from a function and we dont wantto use
	result, _ := div(10, 0)
	fmt.Println(result)


	res, err := divide(10, 0)
	fmt.Println(res,err)
}
