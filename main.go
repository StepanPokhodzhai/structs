package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	stepan := person{
		firstName: "Stepan",
		lastName:  "Pokhodzhai",
		contactInfo: contactInfo{
			email:   "stelkerstepan@gmail.com",
			zipCode: 81160,
		},
	}

	stepan.updateName("Stefan")
	stepan.print()

	/*var stepan person

	stepan.firstName = "Stepan"
	stepan.lastName = "Pokhodzhai"

	fmt.Println(stepan)
	fmt.Printf("%+v", stepan)*/
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
