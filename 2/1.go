package main

import "fmt"

func main() {
	var slice1 []int
	var slice2 []string

	slice1 = append(slice1, 1)
	fmt.Println(slice1)

	slice2 = append(slice2, "a")
	fmt.Println(slice2)

	slice1 = append(slice1, 2, 3)
	fmt.Println(slice1)

	slice2 = append(slice2, "b", "c")
	fmt.Println(slice2)
}
