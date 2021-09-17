package main

import "fmt"

func main() {

	//Strings
	var nameOne string = "Joe"
	var nameTwo = "Luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Peach"
	nameThree = "Luis"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Joshua"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	//Ints
	var ageOne int = 20
	var ageTwo = 25
	ageThree := 15

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory
	var numOne int8 = 30
	var numTwo int8 = -128
	var numThree uint8 = 255

	//Floats
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 38849398499377638484.98
	scoreThree := 2.5
}
