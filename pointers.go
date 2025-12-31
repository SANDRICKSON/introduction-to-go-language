package main

import "fmt"

func pointersExample() {

	//pointers
	//არის ცვლადი რომელიც ინახავს სხვა ცვლადის მისამართს

	var num int = 42
	numPointer := &num //& გამოიყენება რომ მიხვდეს რომ წამოიღოს მისამართი
	fmt.Printf("num is of type %T\n", num)
	fmt.Println(numPointer)  //წერს მისამართს
	fmt.Println(*numPointer) //წერს ამ მისამართზე არსებული ცვლადის მნიშვნელობას
}
