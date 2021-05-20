package main

import "fmt"

type person struct {
	firsName string
	lastName string
	contact  contactinfo
}

type contactinfo struct {
	email string
	zip   int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// func (p *person) updateName(name string) {
// 	p.firsName = name
// }

func (p *person) updateName(name string) {
	// (*p).firsName = name
	(p).firsName = name
}

func main() {

	joe := person{
		firsName: "Watchapon",
		lastName: "Jun-o-pat",
		contact: contactinfo{
			email: "watchapon@igg.com",
			zip:   50200,
		},
	}

	// joePointer := &joe
	fmt.Println(*&*&joe)
	// joePointer.updateName("INW JOE")
	// joe.updateName("INW JOE")
	// joe.print()
}
