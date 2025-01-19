package project2

import (
	"fmt"
	"math/rand"
)

/*
*
continue 基本语法
1.
*/
func continue_test1() {
	//打印随机数，直到出现9
	for {
		intn := rand.Intn(10)
		if intn == 9 {
			fmt.Println("结束本次")
			continue
		}
		if intn == 8 {
			break
		}
		fmt.Printf("随机数：%d\n", intn)
	}
}
