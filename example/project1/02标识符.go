package project01

import "fmt"

/*
*
这一章首先学一下格式化字符串方法

	%v	按值的本来值输出
	%+v	在 %v 基础上，对结构体字段名和值进行展开
	%#v	输出 Go 语言语法格式的值
	%T	输出 Go 语言语法格式的类型和值
	%%	输出 % 本体
	%b	整型以二进制方式显示
	%o	整型以八进制方式显示
	%d	整型以十进制方式显示
	%x	整型以十六进制方式显示
	%X	整型以十六进制、字母大写方式显示
	%U	Unicode 字符
	%f	浮点数
	%p	指针，十六进制方式显示
*/
type people struct {
	height float32
	weight float32
	age    int
	gender string
}

func formatted() {
	var num uint8 = 127
	var identifier = "str"
	var basicPeople people = people{
		height: 182.4,
		weight: 72.3,
		age:    20,
		gender: "男",
	}
	var floatNum = 3.14
	var array = []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("identifier === %v\n", identifier)                  // identifier === str
	fmt.Printf("identifier === %+v\n", basicPeople)                // identifier === {height:182.4 weight:72.3 age:28}
	fmt.Printf("identifier === %#v\n", basicPeople)                // identifier === project01.people{height:182.4, weight:72.3, age:28}
	fmt.Printf("identifier === %T\n", basicPeople)                 // identifier === project01.people
	fmt.Printf("num === %b\n", num)                                // num === 10000000
	fmt.Printf("num === %o\n", num)                                // num === 200
	fmt.Printf("num === %d\n", num)                                // num === 128
	fmt.Printf("num === %x\n", num)                                // num === 80
	fmt.Printf("num === %X, identifier === %X\n", num, identifier) // num === 80, identifier === 737472
	fmt.Printf("numUnicode === %U\n", num)                         // numUnicode === U+0080
	fmt.Printf("num === %f\n", floatNum)                           // num === 3.140000
	fmt.Printf("arrayP === %p\n", array)                           // arrayP === 0xc000010510
}

/*
*
1. 标识符用来命名变量，类型等程序实体。可以使用字母、数字、下划线组成。但是首字母不能是数字
*/
func printIdentifier1() {
	var _i = 1
	var i = "test"

	//var 1i = "test"
	fmt.Printf("_i === %v\n", _i)
	fmt.Printf("i === %v\n", i)

}

/*
*
2. 变量的声明一般是通过var关键字， 可以一次性声明多个变量，并赋值
*/
func printIdentifier2() {
	var identifier1, identifier2 string
	identifier1 = "长江"
	identifier2 = "黄河"
	identifier3, identifier4 := 3, 4
	identifier5 := 3.14
	fmt.Printf("identifier1 = %v, identifier2 = %v\n", identifier1, identifier2)
	fmt.Printf("identifier3 = %v, identifier4 = %v\n", identifier3, identifier4)
	fmt.Printf("identifier5 = %v\n", identifier5)
}
