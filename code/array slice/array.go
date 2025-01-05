package main

import "fmt"

func main() {
	var arr1 [5]string
	arr1[0] = "Subham"
	arr1[1] = "Sobhan"
	fmt.Println("Array elements are:",arr1,",Length is:",len(arr1))
	// %q reprenets quoted string
	fmt.Printf("Array elements are: %q",arr1)

	var arr2 = [5]int{1,2,3,4,5}
	var arr3 = [5]int{1,2,3}
	fmt.Println("Array 2 elements are:",arr2,",Length is:",len(arr2))
	fmt.Println("Array 3 elements are:",arr3,",Length is:",len(arr3))

	 
}