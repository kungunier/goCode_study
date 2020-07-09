package main

import "fmt"

/*
	Go 语言条件语句
	条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。

	Go 语言提供了以下几种条件判断语句：
	语句					描述
	if 语句				if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
	if...else 语句		if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
	if 嵌套语句			你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
	switch 语句			switch 语句用于基于不同条件执行不同动作。
	select 语句			select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。

	注意：Go 没有三目运算符，所以不支持 ?: 形式的条件判断。
*/

/*
	Go 语言 if 语句
	if 语句由布尔表达式后紧跟一个或多个语句组成。
	Go 编程语言中 if 语句的语法如下：
	if 布尔表达式 {
		在布尔表达式为 true 时执行，如果为 false 则不执行
	}
*/

/*
	Go 语言 if...else 语句
	if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
	Go 编程语言中 if...else 语句的语法如下：
	if 布尔表达式 {
		在布尔表达式为 true 时执行
	} else {
		在布尔表达式为 false 时执行
	}
*/

/*
	Go 语言 if 语句嵌套
	可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
	Go 编程语言中 if...else 语句的语法如下：
	if 布尔表达式 1 {
		在布尔表达式 1 为 true 时执行
		if 布尔表达式 2 {
			在布尔表达式 2 为 true 时执行
		}
	}
*/

/*
	Go 语言 switch 语句
	switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。
	switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。
	switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough 。
	Go 编程语言中 switch 语句的语法如下：
	switch var1 {
		case val1:
			...
		case val2:
			...
		default:
			...
	}
	变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。

	Type Switch
	switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
	Type Switch 语法格式如下：
	switch x.(type){
		case type:
			statement(s);
		case type:
			statement(s);
		你可以定义任意个数的case
		default: 可选
			statement(s);
	}

	fallthrough
	使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
*/

/*
	Go 语言 select 语句
	select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。
	select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。
	Go 编程语言中 select 语句的语法如下：
	select {
		case communication clause  :
			statement(s);
		case communication clause  :
			statement(s);
		你可以定义任意数量的 case
		default : 可选
			statement(s);
	}

	以下描述了 select 语句的语法：
	每个 case 都必须是一个通信
	所有 channel 表达式都会被求值
	所有被发送的表达式都会被求值
	如果任意某个通信可以进行，它就执行，其他被忽略。
	如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
	否则：
		如果有 default 子句，则执行该语句。
		如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
*/

func main() {

	//if语句判断大小
	a := 10
	if a < 20 {
		fmt.Println("a小于20")
	}
	fmt.Printf("a的值为：%d\n", a)

	//if...else语句判断大小
	b := 40
	if b < 30 {
		fmt.Println("b小于30")
	} else {
		fmt.Println("b大于30")
	}
	fmt.Printf("b的值为：%d\n", b)

	//Go 语言 if 语句嵌套
	if a != b {
		if a > b {
			fmt.Println("a的值大于b")
		} else {
			fmt.Println("a的值小于b")
		}
	}
	fmt.Printf("a的值为：%d\n", a)
	fmt.Printf("b的值为：%d\n", b)

	//Go 语言 switch 语句
	dengji := "A"
	fenshu := 90
	switch fenshu {
	case 90:
		dengji = "A"
	case 80:
		dengji = "B"
	case 70, 60, 50:
		dengji = "C"
	default:
		dengji = "D"
	}

	switch {
	case dengji == "A":
		fmt.Println("优秀")
	case dengji == "B", dengji == "C":
		fmt.Println("良好")
	case dengji == "D":
		fmt.Println("不及格")
	default:
		fmt.Println("差")
	}
	fmt.Printf("你的等级是：%s\n", dengji)

	//Type Switch
	var c interface{}
	switch d := c.(type) {
	case nil:
		fmt.Printf("c的类型为：%T\n", d)
	case int:
		fmt.Printf("c的类型为：int\n")
	case float32, float64:
		fmt.Printf("c的类型为：float\n")
	case bool, string:
		fmt.Printf("c的类型为：bool或string\n")
	default:
		fmt.Printf("未知型\n")
	}

	//fallthrough
	switch {
	case false:
		fmt.Println("1、case的条件为false")
		fallthrough
	case true:
		fmt.Println("2、case的条件为true")
		fallthrough
	case false:
		fmt.Println("3、case的条件为false")
		fallthrough
	case true:
		fmt.Println("4、case的条件为true")
	case false:
		fmt.Println("5、case的条件为true")
	case true:
		fmt.Println("6、case的条件为true")
		fallthrough
	default:
		fmt.Println("fallthrough用法测试")
	}

	//Go 语言 select 语句
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("接收到", i1, "从c1上")
	case c2 <- i2:
		fmt.Println("发送", i2, "去c2")
	case i3, ok := (<-c3):
		if ok {
			fmt.Println("接收到", i3, "从c3")
		} else {
			fmt.Println("c3已关闭")
		}
	default:
		fmt.Println("体验select的用法")
	}

}
