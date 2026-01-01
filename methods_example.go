package main

import "fmt"

// method არის ფუნქცია, რომელიც გამოიყენება struct / custom type-ზე

// ეს არის კლასის ანალოგი
// go ში არ გვაქ class keyword და არც მემკვიდრეობა
type Author struct {
	name      string // პატარა ასოთი როცა იქმნება ატრიბუტი არის package-private
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
