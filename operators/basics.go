package main

import "fmt"

func main() {
	var num = 1 //assignment operator
	fmt.Println(num == 0 ) //comparison operator
	fmt.Println(num != 0 )//comparison operator
	fmt.Println(num >= 1 )//comparison operator
	fmt.Println(2 / 2)//binary arithmetic operator
	fmt.Println(1 % 2)//binary arithmetic operator
	fmt.Println("a" + "b")//binary arithmetic operator
	num++ //unary operator
	fmt.Println(num)
	fmt.Println(true || false)//boolean operator
	fmt.Println(true && false)//boolean operator
}