package main

import "fmt"

func main() {
	i := 10 // Go는 변수를 정의할 때, :=를 사용한다.

	// Go의 조건문은 괄호를 넣지 않는다.
	if i >= 5 {
		fmt.Println("It's lager then 5")
	}

	// Go의 For문은 괄호를 넣지 않는다.
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
