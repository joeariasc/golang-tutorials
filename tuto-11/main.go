package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":   4.99,
		"pie":    7.99,
		"salad":  6.99,
		"toffee": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//loping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		3113796789: "Jose",
		9930949839: "Mar",
		1039400392: "Jake",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[3113796789])

	phonebook[1039400392] = "Karen"
	fmt.Println(phonebook)
}
