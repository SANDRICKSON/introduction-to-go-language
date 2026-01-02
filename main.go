package main

//ფექიჯის შემოტანა

// ასევე შემიძლია შემოვიტანო _ ით ფექიჯი და რომ არ გამოვიყენოთ არ იქნება ერორი unused import

import (
	"IntroductionToGoLanguage/httpPackages"
	"fmt"
	"net/http"
	_ "strings"
)

func main() {

	//fmt.Println("Hello World")
	//
	//var name = "Sandro"
	//fmt.Println(name)
	//var surname string = "Katamadze"
	//fmt.Println(surname)
	////შეიძლება ასევე ცვლადმა განსაზღვროს ტიპი
	//y := 10
	//fmt.Println(y)
	//
	//// ასევე გვაქვს zero value ტიპებიც
	//// როცა ცვლადს ქმნი და არ აძლევ მნიშვნელობას
	////მაგ:
	//
	//var height float32
	//fmt.Println(height) //0 ს დაბეჭდავს
	//
	//// constants
	//// const არის უცვლელი მნიშვნელობა, რომელიც ფიქსირდება კომპილაციის დროს
	//// const არ იწვევს შეცდომას, თუ არ არის გამოყენებული
	//const Pi = 3.14
	//
	////ჩვენ შეგვიძლია const სია შევქმნათ
	////მაგ: კვირის დღეები
	//
	//const (
	//	Monday    = 1
	//	Tuesday   = 2
	//	Wednesday = 3
	//	Thursday  = 4
	//	Friday    = 5
	//	Saturday  = 6
	//	Sunday    = 7
	//)
	//fmt.Println(Sunday)
	//
	////ასევე მათი გამოყენება შეიძლება არითმეტიკულ და ლოგიკურ ოპერაციებში მაგ:
	//const a = 7
	//const b = 8
	//fmt.Println(a * b)
	//fmt.Println(a / b)
	//fmt.Println(a % b)
	//fmt.Println(a ^ b)
	//
	//const isTrue = true && false
	//fmt.Println(isTrue)
	//
	//// Iota
	//// iota არის სპეციალური კონსტანტა, რომელიც გამოიყენება მხოლოდ const ბლოკში
	//// იწყება 0-დან და ყოველი ახალი ხაზზე იზრდება 1-ით
	//const (
	//	Red = iota
	//	Green
	//	Yellow
	//)
	//
	//fmt.Println(Red, Green, Yellow) // 0 1 2
	//
	////შეგიძლია თავიდანვე მიუთითო სასურველი მნიშვნელობა
	////iota + რასაც უნდა უდრიდეს
	////და იქედან გააგრძელებს ერთით მატებას
	//const (
	//	A = iota + 3
	//	B
	//	C
	//	F
	//)
	//
	////statements
	//
	////if-else
	//
	//var age int = 54
	//if age >= 18 {
	//	fmt.Println("You are an adult")
	//} else {
	//	fmt.Println("You are not an adult")
	//}
	//
	////switch statement
	//
	//day := "Wednesday"
	//switch day {
	//case "Saturday":
	//	fmt.Println("Saturday")
	//case "Monday":
	//	fmt.Println("Monday")
	//case "Tuesday":
	//	fmt.Println("Tuesday")
	//case "Wednesday":
	//	fmt.Println("Wednesday")
	//}
	//
	////looping statements
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	////break:for loop
	//
	//for j := 0; j < 10; j++ {
	//	if j == 5 {
	//		break
	//
	//	}
	//	fmt.Println(j) //0 1 2 3 4
	//	goToExample()
	//	result := add(59, 43)
	//	fmt.Println(result)
	//	deferExample()
	//
	//	//panicExample()
	//	pointersExample()
	//	arrayExample()
	//	sliceExample()
	//
	//}
	//author := Author{
	//	name:      "John",
	//	age:       23,
	//	branch:    "Golang",
	//	particles: "Go Programming Language",
	//	salary:    120,
	//}
	//
	//jumberi := Author{
	//	name:      "Jim",
	//	age:       23,
	//	branch:    "Golang",
	//	particles: "Go Programming Language",
	//}
	//
	//// Method call
	//author.PrintInfo()
	//jumberi.PrintInfo()
	//
	////Error Example
	//errorExample()
	//errorWithPackage()
	//
	//circle := circle{
	//	radius: 40,
	//}
	//fmt.Println(circle.Perimeter())
	//fmt.Println(circle.Area())
	//
	//square := square{
	//	side: 4,
	//}
	//
	//fmt.Println(square.Perimeter())
	//fmt.Println(square.Area())
	//fmt.Println(square.calculateDiagonal())
	//
	//go sayHello()
	//inputExample()
	//
	//mypackage.PrintSandro()

	http.HandleFunc("/hello", httpPackages.HelloHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)

	fileExample()
}
