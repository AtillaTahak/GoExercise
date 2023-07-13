package main

import (
	"fmt"
)

type Singleton struct{
	name []string
}

func (s *Singleton) addName( name string){
	s.name = append(s.name, name)
}
var instance *Singleton

func GetInstance() *Singleton{
	if instance == nil{
		instance = &Singleton{
			name: []string{},
		}

	}
	return instance
}

func (s Singleton) printSingleton(){
	fmt.Println("Singleton")
}

func main(){
	singleton := GetInstance()
	singleton.addName("John")
	singleton.addName("Doe")
	singleton2 := GetInstance()
	singleton2.addName("Jane")
	singleton2.addName("Doe")
	singleton.printSingleton()
	fmt.Println(singleton.name)
}