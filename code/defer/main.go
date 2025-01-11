package main

import "fmt"

func main(){
	fmt.Println("Start")

	//differ will print at the end 
	defer fmt.Println("Middle 1")
	defer fmt.Println("Middle 2")
	defer fmt.Println("Middle 3")

	fmt.Println("End")	
}