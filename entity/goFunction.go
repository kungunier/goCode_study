package main

import (
	"fmt"
	"math"
)

/*
	Go 语言函数
	函数是基本的代码块，用于执行一个任务。
	Go 语言最少有个 main() 函数。
	你可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务。
	函数声明告诉了编译器函数的名称，返回类型，和参数。
	Go 语言标准库提供了多种可动用的内置的函数。例如，len() 函数可以接受不同类型参数并返回该类型的长度。
	如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。
*/

/*
	函数定义
	Go 语言函数定义格式如下：
	func function_name( [parameter list] ) [return_types] {
		函数体
	}

	函数定义解析：
	func：函数由 func 开始声明
	function_name：函数名称，函数名和参数列表一起构成了函数签名。
	parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。
	参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
	return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
	函数体：函数定义的代码集合。
*/

/*
	函数参数
	函数如果使用参数，该变量可称为函数的形参。
	形参就像定义在函数体内的局部变量。
	调用函数，可以通过两种方式来传递参数：
	传递类型		描述
	值传递		值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	引用传递		引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

	默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
*/

/*
	Go 语言函数值传递值
	传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
*/

/*
	Go 语言函数引用传递值
	引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
*/

/*
	函数用法
	函数用法						描述
	函数作为另外一个函数的实参		函数定义后可作为另外一个函数的实参数传入
	闭包							闭包是匿名函数，可在动态编程中使用
	方法							方法就是一个包含了接受者的函数
*/

/*
	Go 语言函数闭包
	Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
*/

/*
	Go 语言函数方法
	Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。
	语法格式如下：
		func (variable_name variable_data_type) function_name() [return_type]{
		函数体
		}
*/

/* 定义结构体 */
type Circle struct {
	radius float64
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func main() {

	//函数调用
	num1 := 100
	num2 := 200
	ret := max(num1, num2)
	fmt.Printf("最大值为:%d\n", ret)

	//函数返回多个值
	x, y := swap("google", "firefox")
	fmt.Println(x, y)

	//Go 语言函数值传递值	函数对参数的修改不影响实际参数的值
	a := 100
	b := 200
	fmt.Printf("交换前a的值为：%d\n", a)
	fmt.Printf("交换前b的值为：%d\n", b)

	exchange1(a, b)
	fmt.Printf("交换后a的值为：%d\n", a)
	fmt.Printf("交换后b的值为：%d\n", b)

	//Go 语言函数引用传递值
	c, d := 100, 200
	fmt.Printf("交换前c的值为：%d\n", c)
	fmt.Printf("交换前d的值为：%d\n", d)

	//	&a 指向 a 指针，a 变量的地址
	//  &b 指向 b 指针，b 变量的地址
	exchange2(&c, &d)
	fmt.Printf("交换后c的值为：%d\n", c)
	fmt.Printf("交换后d的值为：%d\n", d)

	/*
		Go 语言函数作为实参
	*/
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	/*
		Go 语言函数闭包
	*/
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()
	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	/*
		Go 语言函数方法
	*/
	var c1 Circle
	c1.radius = 12.432
	fmt.Println("圆的面积为 = ", c1.getArea())

}

/*
	函数返回两个数的最大值
*/
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/*
	函数返回多个值
*/
func swap(x, y string) (string, string) {
	return y, x
}

/*
	值传递值
*/
func exchange1(a, b int) int {
	var temp int
	temp = a
	a = b
	b = temp
	return temp
}

/*
	引用传递值
*/
func exchange2(c *int, d *int) {
	var temp int
	temp = *c /*保持c地址上的值*/
	*c = *d   /*将d赋值给c*/
	*d = temp /*将temp值赋值给d*/
}

/*
	Go 语言函数闭包
*/
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
