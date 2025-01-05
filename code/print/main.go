package main

import "fmt"

func main(){
	fmt.Println("Hello")

	name := "Subham"
	age := 22
	height := 5.71230

	fmt.Println("name:",name,",age:",age,",height:",height)

	fmt.Printf("name: %s, age: %d, height: %.3f\n",name,age,height)

	fmt.Printf("Type of name is %T \n",name)
	fmt.Printf("Type of age is %T \n",age)
	fmt.Printf("Type of height is %T \n",height)
}