package main

import "fmt"

func main(){
	agesMap := make(map[string]int)
	agesMap["felix"] = 25
	agesMap["victor"] = 12
	fmt.Println(agesMap)
	fmt.Println(agesMap["felix"])
	delete(agesMap, "felix")
	fmt.Println(agesMap)
}