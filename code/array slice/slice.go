package main

import "fmt"

func main() {
	fmt.Println("Slices Example")
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr, len(arr))

	arr = append(arr, 6, 7, 8, 9, 10, 11, 12, 13)
	fmt.Println(arr, len(arr))

	//Empty slice
	arr1 := []string{}
	fmt.Println(arr1, len(arr1))

	// create slice using make()
	arr2 := make([]int, 5)
	fmt.Println(arr2, len(arr2))

	// create slice using make() with capacity
	arr3 := make([]int, 0, 5)
	fmt.Println(arr3, len(arr3), cap(arr3))
}
