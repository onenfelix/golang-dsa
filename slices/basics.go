package main

import "fmt"


func main(){
	var mySlice []string
	mySlice = []string{"First", "Second", "Third"}
	fmt.Println(mySlice)
	emptySlice := make([]string, 3)
	fmt.Println(emptySlice)
	newSlice := append(emptySlice, "Fourth", "Fifth")
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	firstSlice := []string{"First", "Second", "Third"}

	secondSlice := make([]string, 2)

	copy(secondSlice,firstSlice)
	fmt.Println(secondSlice)
}