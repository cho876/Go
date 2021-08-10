package main

import "fmt"

func main() {

	// 기본 for문
	for i := 0; i < 5; i++ {
		fmt.Println("i: ", i) // 0,1,2,3,4
	}

	/* Compile Error
	for i:=0;i<5;i++
	{
		fmt.Println("i: ",i)
	}
	*/

	/* Compile Error
	for i:=0;i<5;i++
		fmt.Println("i: ",i)
	*/

	// for 조건식 (=while문)
	i := 0

	for i < 5 {
		fmt.Println("i: ", i) // 0,1,2,3,4
		i += 1
	}

	// 무한루프
	/*
		for {
			fmt.Println("INF")
		}
	*/

	// 무한루프 break 사용
	j := 0

	for {
		if j > 4 {
			break
		}
		fmt.Println("j: ", j) // 0,1,2,3,4
		j += 1
	}

	// continue 사용
	for k := 0; k < 5; k++ {
		if k == 1 {
			continue
		}
		fmt.Println("k: ", k) // 0,2,3,4
	}

	// 반복문 - 복수개의 변수 사용
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println("i: ", i) // 0,1,2,3,4
		fmt.Println("j: ", j) // 0,2,4,6,8
	}

	/* Compile Error
	for i,j:=0,0; i<5; i++,j+2 {
		fmt.Println("i: ", i)
		fmt.Println("j: ", j)
	}
	*/
}
