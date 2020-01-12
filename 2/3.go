package main

// 추가로 해야함
import (
	"fmt"
)

func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n)
		return s, nil
	}

	return "", fmt.Errorf("%d는 1이 아닙니다.", n)
}
