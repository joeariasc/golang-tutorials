package main

import "fmt"

func main() {
	mybill := newBill("Joe's bill")

	mybill.addItem("onion soup", 8.99)
	mybill.addItem("veg pie", 4.99)
	mybill.addItem("coffee", 3.99)

	mybill.updateTip(3)

	fmt.Println(mybill.format())
}
