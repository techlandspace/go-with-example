package project2

/*
*
goto 基本使用
1. goto可以跳到指定的标签位置，但是只能跳到当前函数的标签位置，不能跳到其他函数的标签位置。
2. goto跳转的位置必须是标签位置，否则编译不通过。
3. 在go程序中一般不推荐使用goto，因为goto会改变程序的执行顺序，导致程序逻辑混乱。
*/
func goto_test1(test int) {
	switch test {
	case 0:
		goto end
	case 1:
		println("我是1")
	default:
		println("我是default")
	}
	println("我是switch中间的内容")
end:
	println("结束咯")
}
