package project01

import "fmt"

/*
*
go语言基础类型由以下几种组成
1.基本类型：

	bool：布尔型，表示真或假。
	int、int8、int16、int32、int64：有符号整型。
	uint、uint8、uint16、uint32、uint64、uintptr：无符号整型。
	float32、float64：浮点型。
	complex64、complex128：复数型。
	byte：别名 uint8，表示字节。
	rune：别名 int32，表示 Unicode 码点。

2.复合类型：

	string：字符串，不可变的字节序列。
	array：固定长度的数组。
	struct：结构体，用户自定义的数据类型，可以包含多个字段。

3.引用类型：

	slice：动态数组，可以改变大小。
	map：键值对集合。
	channel：通道，用于 Goroutine 之间的通信。
	pointer：指针，存储变量的内存地址。
	function：函数类型，可以作为变量或参数传递。
	interface：接口类型，定义了一组方法的集合。

4.特殊类型：

	error：错误类型，通常用于表示错误状态。
*/
func basicDataType() {

	var i byte = 1
	fmt.Printf("i === %d\n", i)
}
