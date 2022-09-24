package main

import "fmt"

func main() {
	var age int = 51

	if age < 18 {
		fmt.Println("He is under age")
	} else if age > 50 {
		fmt.Println("He needs intensive care")
	} else {
		fmt.Println("He is a youth")
	}

	switch age {
	case 14: fmt.Println("He is  small")
	case 1: fmt.Println("He is  sone")
	default: fmt.Println(age)
	}
}