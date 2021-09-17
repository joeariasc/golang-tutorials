package main

import (
	"fmt"
)

func main() {
	age := 29
	name := "Albert"

	// Print
	fmt.Print("Hello, ")
	fmt.Print("World! \n")
	fmt.Print("New Line \n")

	// PrintLn
	fmt.Println("Hello Joe!")
	fmt.Println("Godbay Joe")
	fmt.Println("My age is", age, "and my name is", name)

	// Formatting Strings - Printf %_ = format specifier
	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("your score is %f points! \n", 225.55)
	fmt.Printf("your score is %0.1f points! \n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is: ", str)
}
