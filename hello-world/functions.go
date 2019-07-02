package main

import (
	"fmt"
	"strings"
)

var (
	text = "Hello Im a Function"
)

func converter(stringOne, stringTwo string, intOne int) (int, string, string) { // this line is called the function signature
	stringOne = strings.Title(stringOne)
	stringTwo = strings.ToUpper(stringTwo)
	intOne = intOne + 2

	return intOne, stringOne, stringTwo
}

func main() {

	module := "function basics"
	author := "justin frazier"
	age := 38

	fmt.Println(converter(module, author, age))
}
