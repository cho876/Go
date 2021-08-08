package main

import "fmt"

func main() {

	// 1. 상수
	const INF int = 1e9
	const NAME string = "조성권"

	fmt.Println(INF)
	fmt.Println(NAME)

	// INF = 1	// Compile Error
	// const data int	// Compile Error

	const num1, num2 = 10, 20
	const num3, string1 = 30, "hello"

	// 2. 열거형
	const (
		sunday    = 0
		monday    = 1
		tuesday   = 2
		wednesday = 3
		thursday  = 4
		friday    = 5
		saturday  = 6
	)

	const (
		sunday = iota
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	)

	fmt.Println(sunday)
	fmt.Println(monday)
	fmt.Println(tuesday)
	fmt.Println(wednesday)
	fmt.Println(thursday)
	fmt.Println(saturday)

	// Case 1.
	const (
		a = 1 << iota // a==1 (1 << 0)
		b = 1 << iota // b==2 (1 << 1)
		c = 1 << iota // c==4 (1 << 2)
		d = 1 << iota // d==8 (1 << 3)
	)

	// Case 2.
	const (
		a = iota * 30 // a==0 (0*30)
		b = iota * 30 // b==30 (1*30)
		c = iota * 30 // c==60 (2*30)
		d = iota * 30 // d==90 (3*30)
	)
}
