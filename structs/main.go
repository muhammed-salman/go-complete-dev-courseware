package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	john := person{
		firstName: "John",
		lastName:  "Dsouza",
		contactInfo: contactInfo{
			email:   "john@example.com",
			zipcode: 12345,
		},
	}

	// johnPointer := &john
	// johnPointer.updateName("Johnny")
	//pointer shortcut for ^^ two lines above
	john.updateName("Johnny")

	// fmt.Printf("%+v", john)
	john.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
