package main

import "fmt"

//scope არის ხედვის არე - სად არის ცვლადი ხელმისაწვდომი და ხილვადი

var hobby = "Play Guitar"

func printHobby() {
	height := 93
	fmt.Println(height)
	fmt.Println(hobby)
}

func futToCm() {
	//fmt.Println(height * 0.93) -  არ იმუშავებს რადგან არ არის ხილვადი package ში, ფუნქციაში ჩანს
}
