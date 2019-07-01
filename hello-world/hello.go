package main

import (
	"fmt"
	"os"
	"reflect"
)

// These are global variables set at the "package level"
var (
	globalVar1 = 10
	globalVar2 = 20
	foo        = "bar"
	bar        = "foo"
)

func deReferencing(courseDeRef *string) string {
	*courseDeRef = "Go Fundementals"
	fmt.Println("\nYou are now joining", *courseDeRef)
	return *courseDeRef
}

func variables() {
	const functionVariable string = "I am a function level variable"
	funcVar := 30 // short hand var that must be withing a function and must be initalized when declared
	funcVarPointer := &funcVar
	courseRef := "Web Dev Quest"
	fmt.Println("this is a global var globalVar1 =", globalVar1)
	fmt.Println("You can see the type using the reflect.TypeOf() function", reflect.TypeOf(funcVar))
	fmt.Println("This is the value of funcVar", funcVar)
	fmt.Println("This is the pointer value of funcVar", funcVarPointer)
	fmt.Println("This is a de-refferenced funcVar", *funcVarPointer)
	deReferencing(&courseRef)
	fmt.Println("The course var is now set to", courseRef)
	fmt.Println(`----------Variables Overviews----------
	1. Unanitialized variables get a "zero" value
	2. Package level variables must be declared with a "var"
	3. Package level variables are "global"
	4. Declare function variables with :=
	5. The & symbol references a pointer
	6. The * symbol de-references a pointer
	7. Go supports "type" inferrence
	8. Go passes arguments by "value"`)
}

func envVariables() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}

func main() {
	var a int = 5
	var b int = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a + b)
	fmt.Println(5 + 1)
	fmt.Println("this is go", a, b)
	fmt.Println("\n*********Variables Output**********")
	variables()
	fmt.Println("\n*********envVariables Output**********")
	envVariables()
	fmt.Println(os.Getenv("USER"))
}
