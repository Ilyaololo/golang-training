package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
}

type Person struct {
	FirstName string
	LastName  string
	Car
}

type Car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type Collection struct {
	Persons []Person
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func main() {
	persons := []Person{
		{
			FirstName: "Darth",
			LastName:  "Maul",
			Car: Car{
				Manufacturer: "Toyota",
				Model:        "Camry",
				Doors:        4,
			},
		},
		{
			FirstName: "Darth",
			LastName:  "Vader",
			Car: Car{
				Manufacturer: "Ford",
				Model:        "Mustang",
				Doors:        2,
			},
		},
	}

	data := Collection{
		Persons: persons,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)

	if err != nil {
		log.Fatalln(err)
	}
}
