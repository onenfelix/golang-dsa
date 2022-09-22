package main

import "fmt"

func main() {
	var myArray [3]string = [3]string{"First", "Second", "Third"}
	fmt.Println(myArray)
	fmt.Println(myArray[0])
	myArray[2] = "Another"
	fmt.Println(myArray)
	fmt.Println(len(myArray))
}