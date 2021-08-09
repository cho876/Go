package main

// 패키지 사용
import (
	"fmt"
	"runtime"
)

import "fmt"
import "runtime"


func main() {
	fmt.Println("CPU's count: ", runtime.NumCPU())
}


// 전역 패키지 사용
import . "fmt"

func main() {
	Println("hello world")
}


// 미사용 패키지에 대한 처리
import (
	"fmt"
	_ "runtime"
)

func main() {
	//fmt.Println("hello World") // Compile Error

	fmt.Println("Hello World")
}
