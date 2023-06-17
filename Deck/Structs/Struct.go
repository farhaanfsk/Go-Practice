package main

import "fmt"

func main() {
	person := person{
		firstName: "test",
		lastName:  "test1",
		contact: contact{
			mobile: "12345",
			email:  "mail",
		},
	}
	person.updateName("updated")
	fmt.Println(person)

}

func (p *person) updateName(name string) {
	(*p).firstName = name

}

type person struct {
	firstName string
	lastName  string
	contact   contact
}

type contact struct {
	mobile string
	email  string
}
