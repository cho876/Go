package main

import "fmt"

func main() {

	// 정수 기반, switch문
	i := 5

	switch i {
	case 1:
		fmt.Println("i: 1")
	case 2:
		fmt.Println("i: 2")
	case 3:
		fmt.Println("i: 3")
	case 4:
		fmt.Println("i: 4")
	case 5:
		fmt.Println("i: 5") // return Println
	default:
		fmt.Println("default")
	}

	// 문자열 기반, switch문
	s := "hello"

	switch s {
	case "Hello":
		fmt.Println("Hello")
	case "hellO":
		fmt.Println("hellO")
	case "hello":
		fmt.Println("hello") // return Println
	case "HELLO":
		fmt.Println("HELLO")
	default:
		fmt.Println("Deault")
	}

	// switch문 + fallthrough
	k := 2

	switch k {
	case 1:
		fmt.Println("k: 1")
	case 2:
		fmt.Println("k: 2")
		fallthrough
	case 3:
		fmt.Println("k: 3")
		fallthrough
	case 4:
		fmt.Println("k: 4")
		fallthrough
	default:
		fmt.Println("default")
	}

	// switch문 + 조건문
	j := 2
	str := "world"

	switch j {
	case 1:
		fmt.Println("j: 1")
	case 2:
		if str == "world" {
			fmt.Println("j: 2, str: world") // return Println
		} else {
			fmt.Println("j: 2, str: ", str)
		}
	default:
		fmt.Println("default")
	}

	// 복수개 case문
	h := 3

	switch h {
	case 1, 2:
		fmt.Println("h==1 or h==2")
	case 3, 4:
		fmt.Println("h==3 or h==4") // return Println
	case 5, 6:
		fmt.Println("h==5 or h==6")
	}

	// case 조건문
	a := 3

	switch {
	case a < 3:
		fmt.Println("a < 3")
	case 3 <= a && a < 5:
		fmt.Println("3 <= a < 5") // return Println
	case 6 <= a:
		fmt.Println("6 <= a")
	}
}
