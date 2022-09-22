package main

import (
	"fmt"
	"strings"
)

func main () {
	var name = "Felix Onen, is a software engineer";
	fmt.Println(name);
	fmt.Println(len(name));
	fmt.Println(name[:2]);
	var newstring = name[:];
	fmt.Println(newstring);
	fmt.Println(strings.HasPrefix(name, "Fe"));
	fmt.Println(strings.Count(name, "en"));
	fmt.Printf("%q\n", strings.Split(name, " "))
}