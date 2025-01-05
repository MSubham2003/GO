package main

import "fmt"

func main() {
	// studentGraede := make(map[string]int)

	// studentGraede["Zack"] = 92
	// studentGraede["Jane"] = 96
	// studentGraede["Doe"] = 90
	// studentGraede["Alice"] = 80
	// studentGraede["Bob"] = 93

	studentGraede := map[string]int{
		"Zack":  92,
		"Jane":  96,
		"Doe":   90,
		"Alice": 80,
		"Bob":   93,
	}

	// fmt.Println(studentGraede)

	fmt.Println("Zack's grade is", studentGraede["Zack"])
	fmt.Println("Unknown people grade is", studentGraede["x"])

	// Change value of a key

	studentGraede["Zack"] = 100
	fmt.Println("Zack's grade is", studentGraede["Zack"])

	//Delete a key
	delete(studentGraede, "Zack")
	fmt.Println("Zack's grade is", studentGraede["Zack"])

	// Check if a key exists
	grade, ok := studentGraede["Zack"]
	if ok {
		fmt.Println("Zack's grade is", grade)
	} else {
		fmt.Println("Zack's grade is not found")
	}

	//Iterate in a map
	for key,value:=range studentGraede{
		fmt.Printf("%s's grade is %d\n", key, value)
	}

}
