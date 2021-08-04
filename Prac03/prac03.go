/*
	3장. 변수 선언
*/

package main

import "fmt"

func main() {

	// Case 1 - ':='를 사용하면 var 입력 없이 간단히 변수 선언 가능
	/*
		a, b := 10, 5

		fmt.Println("a: ", a)
		fmt.Println("b: ", b)
	*/

	// Case 2 - 변수의 자료형, 이름을 먼저 선언 후, 값 대입
	/*
		var a, b int
		a = 10
		b = 20

		fmt.Println("a: ", a)
		fmt.Println("b: ", b)
	*/

	// Case 3 - 변수의 선언과 동시에 값 대입
	/*
		var a, b int = 30, 40

		fmt.Println("a: ", a)
		fmt.Println("b: ", b)
	*/

	// Case 4 - 다중 선언

	var (
		a, b int = 10, 20
		c, d     = "Hello", "World"
	)

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)

}
