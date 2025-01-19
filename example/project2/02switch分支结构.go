package project2

import "fmt"

/*
1. switch 语句用于基于不同条件执行不同动作
2. 匹配项后面不需要 break
3. switch 语句可以有多个 case 条件，但是只有一个 case 匹配，那么它将执行匹配的 case 语句，并跳出 switch 语句
4. switch 语句可以有 default 语句，当没有 case 匹配时，它将执行 default 语句
5. case后的表达式可以有多个，使用逗号间隔
6. default 语句可以没有
7. switch fallthrough 允许一个 case 语句执行完，继续执行下一个 case 语句（直接执行，不校验case）
8. type switch: switch 语句可以判断变量的类型
*/
func switch_example(char string) {
	switch char {
	case "apple", "pear":
		fmt.Println("水果")
	case "frag":
		fmt.Println("动物")
	case "computer":
		fmt.Println("机械")
	default:
		fmt.Println("无")
	}
}

func switch_fallthrough_example(char string) {
	switch char {
	case "apple", "pear":
		fmt.Println("水果")
		fallthrough
	case "frag":
		fmt.Println("动物")
		fallthrough
	case "computer":
		fmt.Println("机械")
	default:
		fmt.Println("无")
	}
}

func switch_type_example() {
	var x interface{}
	var y = 1.0
	x = y
	switch i := x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Printf("x的类型： %T\n", i)
	}
}
