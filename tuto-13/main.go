package main

import "fmt"

func updateName(x *string) {
	*x = "wendy"
}

func main() {
	// group A -> strings, ints, bools, floats, arrays, structs
	name := "tifa"

	m := &name

	fmt.Println("memory address m: ", m)
	fmt.Println("value at memory address m: ", *m)

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)
}

/*
|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
| "tifa"  | p0x001  |
|---------|---------|
*/
