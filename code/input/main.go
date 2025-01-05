package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter anything")


	// var value string
	// //Scan can read upto space encounter
	// fmt.Scan(&value)
	// fmt.Println("User entered:", value)

	//this will read input string till we encounter enter
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("User entered:", name)
}
