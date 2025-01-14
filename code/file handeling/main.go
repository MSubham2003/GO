package main

import (
	"fmt"
	"os"
)

func main() {

	// //Create a file
	// file, err := os.Create("test.txt")
	// defer file.Close() //Thile line will free up the file resource after the function ends
	// if err != nil {
	// 	fmt.Println("Error while creating the file",err)
	// 	return
	// }

	// //Write content in file
	// content1 := "Hello World i am in txt file\n"
	// _, error1 := io.WriteString(file, content1)
	// if error1 != nil {
	// 	fmt.Println("Error while writing in file",error1)
	// }
	// content2 := "My name is Subham\n"
	// _, error2 := io.WriteString(file, content2)
	// if error2 != nil {
	// 	fmt.Println("Error while writing in file",error2)
	// }
	// content3 := "Today is Tuesday\n"
	// _, error3 := io.WriteString(file, content3)
	// if error3 != nil {
	// 	fmt.Println("Error while writing in file",error3)
	// }

	// //Open a file
	// file1, err1 := os.Open("test.txt")
	// defer file1.Close()//Thile line will free up the file resource after the function ends
	// if err1 != nil {
	// 	fmt.Println("Error while Opening file",err1)
	// 	return
	// }

	// //Create a buffer to read file content
	// buffer := make([]byte, 1024)

	// //Read file content using buffer
	// for {
	// 	n,err2 := file1.Read(buffer)
	// 	if err2 == io.EOF{
	// 		break
	// 	}
	// 	if err2 != nil{
	// 		fmt.Println("Error while reading file",err2)
	// 		return
	// 	}

	// 	//print file content
	// 	fmt.Print(string(buffer[:n]))
	// }

	//Read the file content using ioutil(now used as os)

	//				It read entire content of file in memory best for small files
	content, err4 := os.ReadFile("test.txt")
	if err4 != nil {
		fmt.Println("Error while reading file", err4)
		return
	}
	//Print file content
	fmt.Println(string(content))

}
