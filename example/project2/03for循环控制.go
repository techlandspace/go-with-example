package project2

import "fmt"

/*
*
for循环四要素
1. 循环变量初始化
2. 循环条件
3. 循环体
4. 循环变量更新
*/
func fortest1() {
	for i := 0; i < 10; i++ {
		println(i)
	}
}

// 死循环
func fortest2() {
	for {
		println("hello world")
	}
}

// for range
func fortest3() {
	str := "hello,world. 中国对世界说话"
	for index, value := range str {
		fmt.Printf("index = %d, value = %c\n", index, value)
	}
}

// 打印金字塔
func forexercise() {
	for i := 0; i <= 6; i++ {
		for j := 0; j < 6-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
