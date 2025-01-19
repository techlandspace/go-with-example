package project2

import "fmt"

/*
*
1.分支控制
  - if else
*/
func processControl1() {

	if true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func processControl2(i int) {
	/**
	golang 支持在if中直接定义一个变量，变量的作用域仅在if中，
	*/
	if num := 1; num < i {
		fmt.Printf("输入参数比条件大\n")
	} else if num == i {
		fmt.Printf("输入参数与条件相等\n")
	} else {
		fmt.Printf("输入参数比条件小\n")
	}
}
