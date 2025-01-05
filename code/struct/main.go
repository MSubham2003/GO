package main

import "fmt"

type student struct {
	Name  string
	Grade int
	Age   int
}

type education struct {
	School    string
	ClassName string
	section   string
}

type parents struct {	
	FatherName string
	MotherName string
} 

type person struct {
	educationdetails education
	parentsdetails parents
}

func main() {
	var students []student
	students = append(students, student{Name: "Zack", Grade: 92, Age: 20})
	students = append(students, student{Name: "Jane", Grade: 96, Age: 21})
	
	fmt.Println(students)


	var student1 student
	student1.Name = "Doe"
	student1.Grade = 90
	student1.Age = 22

	student2:= student{
		Name:  "John",
		Grade: 95,
		Age:   23,
	}

	var student3 = new(student)
	student3.Name = "Dore"
	student3.Grade = 94
	student3.Age = 19
	

	fmt.Println("Name is", student1.Name, "Grade is", student1.Grade, "Age is", student1.Age)
	fmt.Println("Name is", student2.Name, "Grade is", student2.Grade, "Age is", student2.Age)
	fmt.Println("Name is", student3.Name, "Grade is", student3.Grade, "Age is", student3.Age)


	person1 := person{
		educationdetails: education{
			School:    "XYZ",
			ClassName: "10th",
			section:   "A",
		},
		parentsdetails: parents{
			FatherName: "John",
			MotherName: "Jane",
		},
	}

	fmt.Println(person1)
	fmt.Println(person1.educationdetails.School)



}