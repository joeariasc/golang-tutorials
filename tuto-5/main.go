package main

import (
	"fmt"
	"sort"
)

func main() {

	/* greeting := "hello my friends I hope you have a wonderful day"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println("index in", strings.Index(greeting, "wo"))
	fmt.Println(strings.Split(greeting, " "))

	// the original string is unchanged
	fmt.Println("the original string is:", greeting) */

	ages := []int{45, 20, 35, 78, 65, 19}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 78)
	fmt.Println(index)

	names := []string{"Joe", "Luis", "Andres", "Lauren"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Andres"))

}
