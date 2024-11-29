package main

import (
	"fmt"
)

const aConst int = 64

func main() {

	var aString string = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("This variable's type is %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "THis is another string"
	fmt.Println(anotherString)
	fmt.Printf("This variable's type is %T\n", anotherString)

	var anotherInt int = 15
	fmt.Println(anotherInt)
	fmt.Printf("This variable's type is %T\n", anotherInt)

	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("This variable's type is %T\n", myString)

	fmt.Println(aConst)
	fmt.Printf("This variable's type is %T\n", aConst)
}
