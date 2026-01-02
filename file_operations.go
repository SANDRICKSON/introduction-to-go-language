package main

import (
	"bufio" // სწრაფი და უსაფრთხო სკანირების ბიბლიოთეკა
	"fmt"   //კონსოლში დასაბეჭდათ
	"os"    //ოპერაციული სისტემის ფუნქციები(stdin, ფაილები)
)

func fileExample() {
	fmt.Print("Enter some text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(input) // _ ნიშნავს რომ არ გვაინტერესებს რამდენ ბაიტს დაიკავებს ფაილი
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File created successfully")
}
