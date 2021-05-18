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

func main() {

	joe := person{
		firsName: "Watchapon",
		lastName: "Jun-o-pat",
		contact: contactinfo{
			email: "watchapon@igg.com",
			zip:   50200,
		},
	}
	joe.print()

	// joe = person{
	// 	"Watchapon",
	// 	"Jun-o-pat",
	// 	"contact": {
	// 		"watchapon@igg.com",
	// 		50200,
	// 	},
	// }

	// joe.firsName = "Watchapon"
	// joe.lastName = "Jun-o-pat"

}
