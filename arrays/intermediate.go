package main

import "fmt"

func main() {
	var buffer [256]byte //Arrays hold storage for a slice
	var slice []byte = buffer[100:150]
	fmt.Println(slice)
	slice2 := slice[0:10]
	fmt.Println(slice2)
	fmt.Println(&slice2[0] == &buffer[100])
}