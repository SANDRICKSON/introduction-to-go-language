package main

import "fmt"

func arrayExample() {
	numbers := [5]int{1, 2, 3, 4, 5} //[]ში იწერება რამდენელემენტიანი უნდა იყოს მასივი
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	//fmt.Println(numbers[3])
}

// slice არ აქვს შეზღუდვა ზომაზე შეიძლება გაიზარდოს ან შემცირდეს
func sliceExample() {
	names := []string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)
}
