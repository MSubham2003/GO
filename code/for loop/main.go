package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	var numbers []int = []int{1, 2, 3, 4, 5}

	for idx, val := range numbers {
		fmt.Printf("Index %d, Value %d\n", idx, val)
	}

	str := "Hello Word!"

	for index, char := range str {
		fmt.Printf("\nChar at %d is %c", index, char)
	}
}
