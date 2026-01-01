package main

import (
	"errors"
	"fmt"
)

func errorExample() {
	err := fmt.Errorf("This is an error")
	fmt.Println(err)
}

func errorWithPackage() {
	err2 := errors.New("This is an error")
	fmt.Println(err2)
}
