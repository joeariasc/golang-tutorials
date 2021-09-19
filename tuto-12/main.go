package main

import "fmt"

func updateName(x string) string {
	x = "Luis"
	return x
}

func updateMenu(y map[string]float64) {
	y["soup"] = 5.99
}

func main() {
	// group A -> strings, ints, bools, floats, arrays, structs
	name := "Tifany"

	name = updateName(name)

	fmt.Println(name)

	// group A -> slices, maps, functions
	menu := map[string]float64{
		"pie":    7.99,
		"salad":  6.99,
		"toffee": 3.55,
	}

	updateMenu(menu)

	fmt.Println(menu)

}
