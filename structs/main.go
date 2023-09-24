package main

import "fmt"

type contactinfo struct {
	email   string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	contactinfo
}

func main() {

	jim := person{
		firstName: "jim",
		lastName:  "P",
		contactinfo: contactinfo{
			email:   "abc@123",
			pinCode: 670646,
		},
	}

	jim.updateName("jam")
	jim.print()
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
