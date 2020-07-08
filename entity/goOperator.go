package main

import "fmt"

/*
	Go 语言运算符
	运算符用于在程序运行时执行数学或逻辑运算。
	Go 语言内置的运算符有：
		算术运算符
		关系运算符
		逻辑运算符
		位运算符
		赋值运算符
		其他运算符
*/

/*
	算术运算符
	下表列出了所有Go语言的算术运算符。假定 A 值为 10，B 值为 20。
	运算符	描述		实例
	+		相加		A + B 输出结果 30
	-		相减		A - B 输出结果 -10
	*		相乘		A * B 输出结果 200
	/		相除		B / A 输出结果 2
	%		求余		B % A 输出结果 0
	++		自增		A++ 输出结果 11
	--		自减		A-- 输出结果 9
*/

/*
	关系运算符
	下表列出了所有Go语言的关系运算符。假定 A 值为 10，B 值为 20。
	运算符	描述														实例
	==		检查两个值是否相等，如果相等返回 True 否则返回 False。			(A == B) 为 False
	!=		检查两个值是否不相等，如果不相等返回 True 否则返回 False。		(A != B) 为 True
	>		检查左边值是否大于右边值，如果是返回 True 否则返回 False。		(A > B) 为 False
	<		检查左边值是否小于右边值，如果是返回 True 否则返回 False。		(A < B) 为 True
	>=		检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。	(A >= B) 为 False
	<=		检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。	(A <= B) 为 True
*/

/*
	逻辑运算符
	下表列出了所有Go语言的逻辑运算符。假定 A 值为 True，B 值为 False。
	运算符	描述																	实例
	&&		逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。		(A && B) 为 False
	||		逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。		(A || B) 为 True
	!		逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。		!(A && B) 为 True
*/

/*
	位运算符
	位运算符对整数在内存中的二进制位进行操作。
	下表列出了位运算符 &, |, 和 ^ 的计算：
	p		q		p & q		p | q		p ^ q
	0		0		0			0			0
	0		1		0			1			1
	1		1		1			1			0
	1		0		0			1			1

	Go 语言支持的位运算符如下表所示。假定 A 为60，B 为13：
	运算符	描述																	实例
	&		按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。			(A & B) 结果为 12, 二进制为 0000 1100
	|		按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或			(A | B) 结果为 61, 二进制为 0011 1101
	^		按位异或运算符"^"是双目运算符。
			其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。		(A ^ B) 结果为 49, 二进制为 0011 0001
	<<		左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。
			其功能把"<<"左边的运算数的各二进位全部左移若干位，
			由"<<"右边的数指定移动的位数，高位丢弃，低位补0。							A << 2 结果为 240 ，二进制为 1111 0000
	>>		右移运算符">>"是双目运算符。右移n位就是除以2的n次方。
			其功能是把">>"左边的运算数的各二进位全部右移若干位，
			">>"右边的数指定移动的位数。											A >> 2 结果为 15 ，二进制为 0000 1111
*/

/*
	赋值运算符
	下表列出了所有Go语言的赋值运算符。
	运算符	描述											实例
	=		简单的赋值运算符，将一个表达式的值赋给一个左值		C = A + B 将 A + B 表达式结果赋值给 C
	+=		相加后再赋值									C += A 等于 C = C + A
	-=		相减后再赋值									C -= A 等于 C = C - A
	*=		相乘后再赋值									C *= A 等于 C = C * A
	/=		相除后再赋值									C /= A 等于 C = C / A
	%=		求余后再赋值									C %= A 等于 C = C % A
	<<=		左移后赋值									C <<= 2 等于 C = C << 2
	>>=		右移后赋值									C >>= 2 等于 C = C >> 2
	&=		按位与后赋值									C &= 2 等于 C = C & 2
	^=		按位异或后赋值									C ^= 2 等于 C = C ^ 2
	|=		按位或后赋值									C |= 2 等于 C = C | 2
*/

/*
	其他运算符
	下表列出了Go语言的其他运算符。
	运算符	描述					实例
	&		返回变量存储地址		&a; 将给出变量的实际地址。
	*		指针变量				*a; 是一个指针变量
*/

var (
	oper1 = 10
	oper2 = 20
	oper3 = 31
	oper4 int
	oper5 float32
)

func main() {

	//+
	oper4 = oper1 + oper2
	fmt.Println(oper4)
	//-
	oper4 = oper1 - oper2
	fmt.Println(oper4)
	oper4 = oper3 - oper1
	fmt.Println(oper4)
	//*
	oper4 = oper3 * oper2
	fmt.Println(oper4)
	///
	oper5 = float32(oper3 / oper2)
	fmt.Println(oper5)
	//%
	oper4 = oper3 % oper2
	fmt.Println(oper4)
	//++
	oper5++
	fmt.Println(oper5)
	//--
	oper4--
	fmt.Println(oper4)

	a := 10
	b := 20
	//==
	fmt.Println(a == b) //false
	//!=
	fmt.Println(a != b) //true
	//>
	fmt.Println(a > b) //false
	//<
	fmt.Println(a < b) //true
	//>=
	fmt.Println(a >= b) //false
	//<=
	fmt.Println(a <= b) //true

	c := true
	d := false
	//&&
	fmt.Println(c && d) //false
	//||
	fmt.Println(c || d) //true
	//!
	fmt.Println(!(c && d)) //true
	fmt.Println(!(c || d)) //false

	e := 60 /* 60 = 0011 1100 */
	f := 13 /* 13 = 0000 1101 */
	var g int
	//&
	g = e & f /* 12 = 0000 1100 */
	fmt.Println(g)
	//|
	g = e | f /* 61 = 0011 1101 */
	fmt.Println(g)
	//^
	g = e ^ f /* 49 = 0011 0001 */
	fmt.Println(g)
	//<<
	g = e << 2 /* 240 = 1111 0000 */
	fmt.Println(g)
	//>>
	g = e >> 2 /* 15 = 0000 1111 */
	fmt.Println(g)

	var h = 12
	var i *int
	i = new(int)
	*i = h
	//&
	fmt.Println(&h) //0xc0000160c8
	//*
	fmt.Println(*i) //12

}
