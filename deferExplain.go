package main

import (
	"fmt"
)

// defer პირობა სრულდება ფუნქციის შესრულების დასრულების შემდეგ
func deferExample() {
	fmt.Println("Start")
	defer fmt.Println("Middle") //ეს ბოლოს გამოვა
	fmt.Println("End")

}
