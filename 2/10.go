package main

import "fmt"

func main() {
	if 5 > 10 {
		fmt.Println("1")
	} else if 5 == 7 {
		fmt.Println("2")
	} else if 5 == 5 {
		fmt.Println("3")
	} else {
		fmt.Println("4")
	}
}
