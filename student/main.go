package main

import (
	"errors"
	"fmt"
)

type Student struct {
	name  string
	age   int
	grade int
}
const (
	PrintAllStudents = iota + 1
	ShowHighestGrade
	ShowLowestGrade
	CreateNewStudent
	SearchStudent
	Exit
)
func createStudent() Student {
	var name string
	var age, grade int

	fmt.Print("Enter name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("Invalid input for name")
		return Student{"", 0, 0}
	}

	fmt.Print("Enter age: ")
	_, err = fmt.Scan(&age)
	if err != nil {
		fmt.Println("Invalid input for age")
		return Student{"", 0, 0}
	}

	fmt.Print("Enter grade: ")
	_, err = fmt.Scan(&grade)
	if err != nil {
		fmt.Println("Invalid input for grade")
		return Student{"", 0, 0}
	}

	return Student{name, age, grade}
}


func (s Student) find (students []Student, name string) (Student,error) {
	for _, student := range students {
		if student.name == name {
			return student, nil
		}
	}
	return Student{"", 0, 0}, errors.New("Student not found")
}
func main() {
	students := []Student{
		{"John", 21, 90},
		{"Jane", 21, 95},
		{"Bob", 21, 92},
	}
	for chosen := true; chosen; {
		fmt.Println("1. Print all students")
		fmt.Println("2. Print student with highest grade")
		fmt.Println("3. Print student with lowest grade")
		fmt.Println("4. Create new student")
		fmt.Println("5. Search for student with name")
		fmt.Println("6. Exit")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case PrintAllStudents:
			fmt.Println("1. Print all students")
			for _, student := range students {
				fmt.Println(student)
			}
		case ShowHighestGrade:
			fmt.Println("2. Print student with highest grade")
			highestGrade := 0
			for _, student := range students {
				if student.grade > highestGrade {
					highestGrade = student.grade
				}
			}
			for _, student := range students {
				if student.grade == highestGrade {
					fmt.Println(student)
				}
			}
		case ShowLowestGrade:
			fmt.Println("3. Print student with lowest grade")
			lowestGrade := 100
			for _, student := range students {
				if student.grade < lowestGrade {
					lowestGrade = student.grade
				}
			}
			for _, student := range students {
				if student.grade == lowestGrade {
					fmt.Println(student)
				}
			}
		case CreateNewStudent:
			fmt.Println("4. Create new student")
			student := createStudent()
			students = append(students, student)
		case SearchStudent:
			fmt.Println("5. Search for student with name")
			var name string
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			student := Student{}
			result,err := student.find(students, name)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		case Exit:
			fmt.Println("6. Exit")
			chosen = false
		default:
			fmt.Println("Invalid choice")
		}
	}

}
