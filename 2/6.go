package main

import "fmt"

func main() {
	map1 := make(map[string]string)
	map1["사과"] = "Apple"
	map1["바나나"] = "Banana"
	map1["오렌지"] = "Orange"
	map1["딸기"] = "Strawberry"
	map1["수박"] = "Watermelon"

	fmt.Println(map1)
}
