package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// 문자열 생성
	var string1 string = "조성권"
	// string1 := "조성권"

	var string2 string = "Hello"

	fmt.Println(len(string1))                    // 9 : UTF-8 인코딩의 Byte 길이
	fmt.Println(len(string2))                    // 5: 알파벳 5글자
	fmt.Println(utf8.RuneCountInString(string1)) // 2: 실제 한글의 갯수 기

	// 문자열 연산
	var string3 string = "Hello"
	var string4 string = "World"

	string5 := string3 + string4

	/*
		var string5 string
		string5 = string3 + string4
	*/

	fmt.Println("string5: ", string5)          // HelloWorld
	fmt.Println("string5[0]: ", string5[0])    // 72: H의 10진수 값
	fmt.Printf("string5[0]: %c\n", string5[0]) // H: 실제 문자

	var string6 string = "Hello"
	var string7 string = "world"
	var string8 string = "hello"
	var string9 string = "Hello"

	fmt.Println("string6 == String7? ", string6 == string7) // false
	fmt.Println("string6 == string8? ", string6 == string8) // false
	fmt.Println("string6 == string9? ", string6 == string9) // true

	// 문자열 수정
	var string10 string = "Hello World"
	string10 = "Hi I'm Sungkwon"

	fmt.Println("string10: ", string10)

	//string10[0] = 'S'	// Compile Error - 문자열의 요소 수정 불

}
