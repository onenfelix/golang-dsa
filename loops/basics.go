package main

import "fmt"

func main(){
	for i := 0; i < 5; i++ {
		fmt.Printf("current iteration %d value \n", i);
	}

	i := 0

	for i < 2 {
		fmt.Println(i)
		i++
	}

	numbers := [3]int{1,2,3}

	for i, num := range numbers {
		fmt.Printf("%d:%d\n", i, num);
	}
}