package main

import "fmt"

func panicExample() {
	if 95%5 == 0 {
		panic("Error") //panic არის ქივორდი რომელიც გაისვრის exceptionად ჩემს მიერ დაწერილ ტექსტს თუ პირობა სწორია
	}
}

func goToExample() {
	k := 0
start:
	if k < 5 {
		fmt.Println(k) //0,1,2,3,4

		k++
		goto start
	}
}
