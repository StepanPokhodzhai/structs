package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	stepan := person{firstName: "Stepan", lastName: "Pokhodzhai"}
	fmt.Println(stepan)
}
