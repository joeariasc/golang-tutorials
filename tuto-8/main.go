package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning %v \n", name)
}

func sayBay(name string) {
	fmt.Printf("Goodbye %v \n", name)
}

func cycleNames(names []string, f func(string)) {
	for _, value := range names {
		f(value)
	}
}

func cycleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	/* sayGreeting("Joe")
	sayGreeting("Norman")
	sayBay("Joe") */
	names := []string{"Joe", "Luis", "Andres", "Lauren", "Lilia"}
	cycleNames(names, sayGreeting)
	cycleNames(names, sayBay)

	a1 := cycleArea(10.5)
	a2 := cycleArea(15)

	fmt.Println(a1, a2)

}
