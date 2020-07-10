package main

import "fmt"

/*
	Go 语言数组
	Go 语言提供了数组类型的数据结构。
	数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。
	相对于去声明 number0, number1, ..., number99 的变量，使用数组形式 numbers[0], numbers[1] ..., numbers[99] 更加方便且易于扩展。
	数组元素可以通过索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。
*/

/*
	声明数组
	Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：
	var variable_name [SIZE] variable_type
*/

/*
	初始化数组
	初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
	如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小
*/

/*
	访问数组元素
	数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。
*/

/*
	更多内容
	内容				描述
	多维数组			Go 语言支持多维数组，最简单的多维数组是二维数组
	向函数传递数组		你可以向函数传递数组参数
*/

/*
	Go 语言多维数组
	Go 语言支持多维数组，以下为常用的多维数组声明方式：
	var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

	二维数组
	二维数组是最简单的多维数组，二维数组本质上是由一维数组组成的。二维数组定义方式如下：
	var arrayName [ x ][ y ] variable_type
*/

/*
	Go 语言向函数传递数组
	如果你想向函数传递数组参数，你需要在函数定义时，声明形参为数组，我们可以通过以下两种方式来声明：

	方式一
	形参设定数组大小：

	void myFunction(param [10]int)
	{
	.
	.
	.
	}

	方式二
	形参未设定数组大小：

	void myFunction(param []int)
	{
	.
	.
	.
	}
*/

func main() {

	/* 声明数组 */
	var balance1 [10]float32
	/* 初始化数组 */
	balance1 = [10]float32{100.12, 23.432, 321.2, 234.0}

	/* 声明数组，不指定数组长度 */
	var balance2 = [...]float32{123, 432.32, 534.0, 432, 543.234}

	fmt.Println("balance1[4] = ", balance1[4])
	fmt.Println("balance2[4] = ", balance2[4])

	/* 访问数组元素 */
	floatNum1 := balance2[3]
	fmt.Println(floatNum1)

	/* 数组完整操作（声明、赋值、访问） */
	var n [10]int
	var i, j int
	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	/* 输出数组每个元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("第 %d 个值为：%d\n", j+1, n[j])
	}

	/*
		初始化二维数组
	*/
	var twoArray1 = [3][4]int{
		{1, 2, 3, 4},
		{11, 22, 33, 44},
		{111, 222, 333, 444},
	}
	//访问二维数组
	val1 := twoArray1[1][3]
	fmt.Println(val1)

	//二维数组可以使用循环嵌套来输出元素
	/* 数组5行2列 */
	var a = [5][2]int{
		{0, 0},
		{1, 2},
		{2, 4},
		{3, 6},
		{4, 8},
	}
	var i1, j1 int
	for i1 = 0; i1 < 5; i1++ {
		for j1 = 0; j1 < 2; j1++ {
			fmt.Printf("a[%d][%d] = %d\n", i1, j1, a[i1][j1])
		}
	}

}
