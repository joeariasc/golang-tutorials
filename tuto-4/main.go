package main

import "fmt"

func main() {
	//var ages [3]int = [3]int{20,25,13}
	var ages = [3]int{20, 25, 13}
	fmt.Println(ages, len(ages))

	names := [4]string{"Manuel", "Jose", "Angel", "Maria"}
	names[1] = "Mario"
	fmt.Println(names, len(names))

	//Slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	//Slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "Koopa")
	fmt.Println(rangeOne)
}
