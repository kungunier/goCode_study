package main

import (
	"fmt"
	"unsafe"
)

/*
	Go 语言常量
	常量是一个简单值的标识符，在程序运行时，不会被修改的量。
	常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
*/
/*
	常量的定义格式：const 常量名 = value
	多个相同类型的声明可以简写为：const 常量名1, 常量名2 = value1, value2
*/
/*
	常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过
*/

const (
	cl1 = "常量1"
	cl2 = len(cl1)
	cl3 = unsafe.Sizeof(cl1)
)

/*
	iota
	iota，特殊常量，可以认为是一个可以被编译器修改的常量。
	iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
	iota 只是在同一个 const 常量组内递增，每当有新的 const 关键字时，iota 计数会重新开始
*/

const (
	a = iota       //0
	b              //1
	c              //2
	d = "验证iota用法" //验证iota用法，iota += 1
	e              //验证iota用法，iota += 1
	f              //验证iota用法，iota += 1
	g = 1000       //1000，iota += 1
	h              //1000，iota += 1
	i = 10         //10，iota += 1
	j = iota       //9
	k              //10
	l              //11
)

const a1 = iota
const b1 = iota

/*
	iota的位移
	左移运算符 << 是双目运算符。左移 n 位就是乘以 2 的 n 次方。 其功能把 << 左边的运算数的各二进位全部左移若干位，由 << 右边的数指定移动的位数，高位丢弃，低位补 0。
	右移运算符 >> 是双目运算符。右移 n 位就是除以 2 的 n 次方。 其功能是把 >> 左边的运算数的各二进位全部右移若干位， >> 右边的数指定移动的位数。
*/

const (
	m = 1 << iota
	n = 3 << iota
	o
	p
)

/*
	在定义常量组时，如果不提供初始值，则表示将使用上行的表达式
*/

const (
	cla = 1
	clb
	clc
	cld
)

func main() {

	//测试定义的常量使用
	fmt.Println(cl1, cl2, cl3)

	//测试iota特殊常量的用法
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l)
	fmt.Println(a1, b1) //iota 只是在同一个 const 常量组内递增，每当有新的 const 关键字时，iota 计数会重新开始

	//测试iota的位移
	fmt.Println(m) //m = 1<<0	1向左位移0位，不变仍为1
	fmt.Println(n) //n = 3<<1	3向左位移1位，变为二进制110，即6
	fmt.Println(o) //o = 3<<2	3向左位移2位，变为二进制1100，即12
	fmt.Println(p) //p = 3<<3	3向左位移3位，变为二进制11000，即24

	//不提供初始值，则表示将使用上行的表达式
	fmt.Println(cla, clb, clc, cld)

}
