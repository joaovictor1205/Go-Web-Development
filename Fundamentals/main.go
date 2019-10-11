package main

import (
	"fmt"
	"encoding/json"
)

type person struct{
	F_name 	string	`json:"first"`
	L_name 	string	`json:"last"`
}

type student struct{
	person
	Studying bool	`json:"studying"`
}

func (p person) create(){
	fmt.Println("Creating person")
}

type human interface{
	create()
}

func saySomething(h human){
	h.create()
}

func main(){

	personOne := person {
		F_name: "Person",
		L_name: "One",
	}

	personOneJSON, _ := json.Marshal(personOne)
	fmt.Println(string(personOneJSON))


	studentOne := student{
		person: person{
			F_name: "Student",
			L_name:"One",
		},
		Studying: true,
	}

	studentOneJSON, _ := json.Marshal(studentOne)
	fmt.Println(string(studentOneJSON))

	saySomething(personOne)
	saySomething(studentOne)
	
}