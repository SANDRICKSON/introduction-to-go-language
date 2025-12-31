package main

import "fmt"

// method არის ფუნქცია, რომელიც გამოიყენება struct / custom type-ზე
type Author struct {
	name      string
	age       int
	branch    string
	particles string
	salary    int
}

// Method
func (author Author) PrintInfo() {
	fmt.Println("Name:", author.name)
	fmt.Println("Age:", author.age)
	fmt.Println("Branch:", author.branch)
	fmt.Println("Particles:", author.particles)
	fmt.Println("Salary:", author.salary)
}
