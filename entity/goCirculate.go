package main

import "fmt"

/*
	Go 语言循环语句

	Go 语言提供了以下几种类型循环处理语句：
	循环类型		描述
	for 循环		重复执行语句块
	循环嵌套		在 for 循环中嵌套一个或多个 for 循环
*/

/*
	Go 语言 for 循环
	for 循环是一个循环控制结构，可以执行指定次数的循环。

	Go 语言的 For 循环有 3 种形式，只有其中的一种使用分号。
	和 C 语言的 for 一样：
	for init; condition; post { }
	和 C 的 while 一样：
	for condition { }
	和 C 的 for(;;) 一样：
	for { }

	init： 一般为赋值表达式，给控制变量赋初值；
	condition： 关系表达式或逻辑表达式，循环控制条件；
	post： 一般为赋值表达式，给控制变量增量或减量。
	for语句执行过程如下：
		1、先对表达式 1 赋初值；
		2、判别赋值表达式 init 是否满足给定条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition；
		否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。

	for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
	for key, value := range oldMap {
		newMap[key] = value
	}
*/

/*
	Go 语言循环嵌套
	以下为 Go 语言嵌套循环的格式：
	for [condition |  ( init; condition; increment ) | Range]
	{
		for [condition |  ( init; condition; increment ) | Range]
		{
			statement(s);
		}
		statement(s);
	}
*/

/*
	循环控制语句
	循环控制语句可以控制循环体内语句的执行过程。

	GO 语言支持以下几种循环控制语句：
	控制语句			描述
	break 语句		经常用于中断当前 for 循环或跳出 switch 语句
	continue 语句	跳过当前循环的剩余语句，然后继续进行下一轮循环。
	goto 语句		将控制转移到被标记的语句。
*/

/*
	Go 语言 break 语句
	Go 语言中 break 语句用于以下两方面：
		用于循环语句中跳出循环，并开始执行循环之后的语句。
		break 在 switch（开关语句）中在执行一条 case 后跳出语句的作用。
		在多重循环中，可以用标号 label 标出想 break 的循环。
*/

/*
	Go 语言 continue 语句
	Go 语言的 continue 语句 有点像 break 语句。但是 continue 不是跳出循环，而是跳过当前循环执行下一次循环语句。
	for 循环中，执行 continue 语句会触发 for 增量语句的执行。
	在多重循环中，可以用标号 label 标出想 continue 的循环。
*/

/*
	Go 语言 goto 语句
	Go 语言的 goto 语句可以无条件地转移到过程中指定的行。
	goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
	但是，在结构化程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。
	goto 语法格式如下：
	goto label;
	..
	.
	label: statement;
*/

func main() {

	//for循环示例——计算1到10的总和
	sum := 0
	for i := 0; i <= 10; i++ {
		sum = sum + i
	}
	fmt.Printf("1到10的总和为%d\n", sum)

	//init 和 post 参数是可选的，我们可以直接省略它，类似 While 语句。
	sum = 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Printf("1到10的总和为%d\n", sum)
	//另一种写法
	sum = 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Printf("1到10的总和为%d\n", sum)

	//For-each range 循环
	//这种格式的循环可以对字符串、数组、切片等进行迭代输出元素。
	strings := []string{"数据1", "数据2"}
	for a, b := range strings {
		fmt.Println(a, b)
	}

	numbers := [6]int{1, 2, 3, 4}
	for a, b := range numbers {
		fmt.Printf("第 %d 位值为 %d\n", a, b)
	}

	//使用循环嵌套来输出 2 到 100 间的素数
	//定义局部变量
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= i/j; j++ {
			if i%j == 0 {
				break //如果发现因子，则不是素数
			}
		}
		if j > i/j {
			fmt.Printf("%d  是素数\n", i)
		}
	}

	//九九乘法表
	for m := 1; m < 10; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%dx%d=%d   ", n, m, m*n)
		}
		fmt.Println("")
	}

	//break跳出循环
	var a = 10
	for a < 20 {
		fmt.Printf("a的值为：%d\n", a)
		a++
		if a > 15 {
			break
		}
	}

	//在多重循环中，可以用标号 label 标出想 break 的循环。
	//不使用标记
	fmt.Println("----break----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i:%d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2:%d\n", i2)
			break
		}
	}
	//使用标记
	fmt.Println("----break label----") //标记的名字可以随意定义	Manchu
Manchu:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i:%d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2:%d\n", i2)
			break Manchu
		}
	}

	//Go 语言 continue 语句
	var b = 10
	for b < 20 {
		if b == 15 {
			b = b + 1
			continue //跳出此次循环
		}
		fmt.Printf("b的值为：%d\n", b)
		b++
	}

	//在多重循环中，可以用标号 label 标出想 continue 的循环。
	//不使用标记
	fmt.Println("----continue----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i:%d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2:%d\n", i2)
			continue
		}
	}
	//使用标记
rater:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i:%d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2:%d\n", i2)
			continue rater
		}
	}

	//Go 语言 goto 语句
	var c = 10
loop:
	for c < 20 {
		if c == 15 {
			c = c + 1
			goto loop
		}
		fmt.Printf("c的值为：%d\n", c)
		c++
	}

}
