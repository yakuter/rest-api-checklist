package main

import (
	"fmt"
)

func recoverSetName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func setName(name *string) {

	if name == nil {
		panic("runtime error: first name cannot be nil")
	}
}

func main() {
	defer recoverSetName()
	setName(nil)
}
