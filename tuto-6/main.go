package main

import "fmt"

func main() {
	/* x := 0

	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	} */

	/* for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	} */

	names := []string{"Joe", "Luis", "Andres", "Lauren"}

	/* for i := 0; i < len(names); i++ {
		fmt.Printf("the name in the position %v is %v \n", i, names[i])
	} */

	/* 	for index, value := range names {
		fmt.Printf("the name in the position %v is %v \n", index, value)
	} */

	for _, value := range names {
		fmt.Printf("the name is %v \n", value)
		value = "new string"
	}

	fmt.Println(names)
}
