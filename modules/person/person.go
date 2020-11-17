package main

import "fmt"

type Contact struct {
	Email string
	Zip   int
}

type Person struct {
	FirstName string
	LastName  string
	Contact
}

func (p *Person) SetFirstName(firstName string) {
	p.FirstName = firstName
}

func (p Person) Print() {
	fmt.Printf("%+v", p)
}

func init() {
	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Contact: Contact{
			Email: "example@email.com",
			Zip:   111222,
		},
	}

	p.Print()
}
