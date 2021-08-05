package main

import (
	"fmt"
)

func main() {

	// 1. 정수 Type 예시
	var num3 uint16 = 10
	var num4 uint32 = 80000

	fmt.Println(uint16(num4))        // 14464
	fmt.Println(num3 + uint16(num4)) // 14474

	// 2. 실수 Type 예시
	var data float64 = 10.0

	for i := 0; i < 10; i++ {
		data -= 0.1
	}

	fmt.Println(data) // 9.000000000000004

	const epsilon = 1e-14

	if math.Abs(data-9.0) <= epsilon {
		fmt.Println("It's smaller than epsilon")
	} else {
		fmt.Println("It's bigger than epsilon")
	}

	// 3.정수형과 실수형의 연산
	var num5 int = 10
	var num6 float32 = 2.2

	fmt.Println(num5 + num6)          // Compile Error
	fmt.Println(num5 + int(num6))     // 10 + 2 = 12
	fmt.Println(float32(num5) + num6) // 10.0 + 2.2 = 12.2

}
