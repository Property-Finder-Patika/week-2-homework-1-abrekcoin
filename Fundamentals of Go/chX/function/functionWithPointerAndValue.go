package main

import "fmt"

func main() {
	p := Person{"Selim", 22}
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
	updatePersonValue(p)
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
	updatePersonPointer(&p)
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
}

func updatePersonValue(p Person) {
	p.name = "New Name"
	p.age++
}

func updatePersonPointer(p *Person) {
	p.name = "New Name"
	p.age++
}

type Person struct {
	name string
	age  int
}
