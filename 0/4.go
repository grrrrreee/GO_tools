package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(2,3))
}


package main
 
import "fmt"
 
func add_sub(x int, y int) (int, int) {
   return x + y, x-y
}
 
func main() {
   fmt.Println(add(2,3))
}
