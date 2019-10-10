package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	F_name string `json:"first"`
	L_name string `json:"last"`
}

type student struct {
	person
	Studying bool `json:"studying"`
}

func (p person) create() {
	fmt.Println("Creating person")
}

func main() {

	personOne := person{
		F_name: "Person",
		L_name: "One",
	}
	personOne.create()

	personOneJSON, _ := json.Marshal(personOne)
	fmt.Println(string(personOneJSON))

	studentOne := student{
		person: person{
			F_name: "Student",
			L_name: "One",
		},
		Studying: true,
	}
	studentOne.create()

	studentOneJSON, _ := json.Marshal(studentOne)
	fmt.Println(string(studentOneJSON))

}
