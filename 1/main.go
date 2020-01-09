package main

import (
	"fmt"
)

func main() {
	fmt.Println(1 + 2)
	fmt.Println(2 - 1)
	fmt.Println(2 * 2)
	fmt.Println(6 / 2)
	fmt.Println(7 % 3)
}

// go routine example 2

// func main() {
// 	fmt.Println("main 함수 시작", time.Now())

// 	go long()
// 	go short()

// 	time.Sleep(5 * time.Second) // 5초 대기
// 	fmt.Println("main 함수 종료", time.Now())
// }

// func long() {
// 	fmt.Println("long 함수 시작", time.Now())
// 	time.Sleep(3 * time.Second) // 3초 대기
// 	fmt.Println("long 함수 종료", time.Now())
// }

// func short() {
// 	fmt.Println("short 함수 시작", time.Now())
// 	time.Sleep(1 * time.Second) // 1초 대기
// 	fmt.Println("short 함수 종료", time.Now())
// }
