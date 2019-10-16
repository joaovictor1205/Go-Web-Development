package main

import (
	"log"
	"os"
	"text/template"
)

func main(){

	tpl, err := template.ParseFiles("tpl.gohtml")

	if err != nil {
		log.Fatal(err)
	}

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}

	defer newFile.Close()

	err = tpl.Execute(newFile, nil)
	if err != nil{
		log.Fatal(err)
	}
}