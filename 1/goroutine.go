package main

import (
	"fmt"
	"time"
)

func A() {
	for i := 0; i < 10; i++ {
		fmt.Println("A : ", i)
	}
}

func B() {
	for i := 10; i < 20; i++ {
		fmt.Println("B : ", i)
	}
}

func main() {
	// 1. goroutine을 사용하였을 때
	// go A()
	// go B()
	// 2. goroutine을 사용하지 않았을 때
	A()
	B()
	time.Sleep(5 * time.Second)
	fmt.Println("goroutine 사용안함")
	fmt.Println("main 함수 종료", time.Now())
}
