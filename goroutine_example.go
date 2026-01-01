package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello World")
		time.Sleep(time.Millisecond * 500)
	}
}
