package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name  string
	age   int
	grade int
}

func createStudent() Student {
	var name string
	var age int
	var grade int
	fmt.Print("Enter name: ")
	fmt.Scan(&name)
	fmt.Print("Enter age: ")
	fmt.Scan(&age)
	fmt.Print("Enter grade: ")
	fmt.Scan(&grade)
	if reflect.TypeOf(name).Kind() != reflect.String {
		fmt.Println("Invalid name")
		return Student{"", 0, 0}
	}
	if reflect.TypeOf(age).Kind() != reflect.Int {
		fmt.Println("Invalid age")
		return Student{"", 0, 0}
	}
	if reflect.TypeOf(grade).Kind() != reflect.Int {
		fmt.Println("Invalid grade")
		return Student{"", 0, 0}
	}
	return Student{name, age, grade}
}

func (s Student) find (students []Student, name string) Student {
	for _, student := range students {
		if student.name == name {
			return student
		}
	}
	return Student{"", 0, 0}
}
func main() {
	students := make([]Student, 0)
	students = append(students, Student{"John", 21, 90})
	students = append(students, Student{"Jane", 21, 95})
	students = append(students, Student{"Bob", 21, 92})
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
		case 1:
			fmt.Println("1. Print all students")
			for _, student := range students {
				fmt.Println(student)
			}
		case 2:
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
		case 3:
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
		case 4:
			fmt.Println("4. Create new student")
			student := createStudent()
			students = append(students, student)
		case 5:
			fmt.Println("5. Search for student with name")
			var name string
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			student := Student{}
			result := student.find(students, name)
			if result.name == "" {
				fmt.Println("Student not found")
			} else {
				fmt.Println(result)
			}
		case 6:
			fmt.Println("6. Exit")
			chosen = false
		default:
			fmt.Println("Invalid choice")
		}
	}

}
