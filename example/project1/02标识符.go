package main

import "fmt"

/**
1.标识符用来命名变量，类型等程序实体。可以使用字母、数字、下划线组成。但是首字母不能是数字
*/

func main() {
	var _i = 1
	var i = "test"
	//var 1i = "test"
	fmt.Printf("_i === %v", _i)
	fmt.Printf("i === %v", i)
}
