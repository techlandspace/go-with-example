package project2

import (
	"fmt"
	"math/rand"
)

/*
*
break 基本语法
1. break 语句出现在单层循环中，直接跳出循环
2. break 语句出现在多层循环中，可以通过指定label来指定跳出哪个循环
*/
func break_test1() {
	//打印随机数，直到出现9
	for {
		intn := rand.Intn(10)
		fmt.Printf("随机数：%d\n", intn)
		if intn == 9 {
			break
		}
	}
}

func break_test2() {
label1:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i=%d,j=%d\n", i, j)
			if j == 2 {
				break label1
			}
		}
	}
}
