package main

import "fmt"

/*
	Go 语言简介
	Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。
	Go是从2007年末由Robert Griesemer, Rob Pike, Ken Thompson主持开发，后来还加入了Ian Lance Taylor, Russ Cox等人，
	并最终于2009年11月开源，在2012年早些时候发布了Go 1稳定版本。现在Go的开发已经是完全开放的，并且拥有一个活跃的社区。
*/
/*
	Go 语言特色
	简洁、快速、安全
	并行、有趣、开源
	内存管理、数组安全、编译迅速
*/
/*
	Go 语言用途
	Go 语言被设计成一门应用于搭载 Web 服务器，存储集群或类似用途的巨型中央服务器的系统编程语言。
	对于高性能分布式系统领域而言，Go 语言无疑比大多数其它语言有着更高的开发效率。它提供了海量并行的支持，这对于游戏服务端的开发而言是再好不过了。
*/
/*
	第一行代码 package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，
	如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。

	下一行 import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。

	下一行 func main() 是程序开始执行的函数。
	main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。

	下一行是注释，在程序执行时将被忽略。单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。
	多行注释也叫块注释，均已以 /* 开头，并以 * /结尾，且不可以嵌套使用，多行注释一般用于包的文档描述或注释成块的代码片段。

	下一行 fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。
	使用 fmt.Print("hello, world\n") 可以得到相同的结果。
	Print 和 Println 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。

	当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，
	那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；
	标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。
*/
/*
	关键字
	下面列举了 Go 代码中会使用到的 25 个关键字或保留字：
	break		default			func		interface		select
	case		defer			go			map	struct
	chan		else			goto		package			switch
	const		fallthrough		if			range			type
	continue	for	import		return		var
*/
/*
	除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：
	append	bool	byte	cap			close	complex	complex64	complex128	uint16
	copy	false	float32	float64		imag	int		int8		int16		uint32
	int32	int64	iota	len			make	new		nil			panic		uint64
	print	println	real	recover		string	true	uint		uint8		uintptr
*/
/*
	Go 语言数据类型
	在 Go 编程语言中，数据类型用于声明函数和变量。
	数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候才需要申请大内存，就可以充分利用内存。
	Go 语言按类别有以下几种数据类型：
	序号	类型和描述
	1	布尔型：布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。
	2	数字类型：整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。
	3	字符串类型：字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。
	4	派生类型：包括：
				(a) 指针类型（Pointer）
				(b) 数组类型
				(c) 结构化类型(struct)
				(d) Channel 类型
				(e) 函数类型
				(f) 切片类型
				(g) 接口类型（interface）
				(h) Map 类型
*/
/*
	数字类型
	Go 也有基于架构的类型，例如：int、uint 和 uintptr。
	序号	类型和描述
	1	uint8：无符号 8 位整型 (0 到 255)
	2	uint16：无符号 16 位整型 (0 到 65535)
	3	uint32：无符号 32 位整型 (0 到 4294967295)
	4	uint64：无符号 64 位整型 (0 到 18446744073709551615)
	5	int8：有符号 8 位整型 (-128 到 127)
	6	int16：有符号 16 位整型 (-32768 到 32767)
	7	int32：有符号 32 位整型 (-2147483648 到 2147483647)
	8	int64：有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
*/
/*
	浮点型
	序号	类型和描述
	1	float32：IEEE-754 32位浮点型数
	2	float64：IEEE-754 64位浮点型数
	3	complex64：32 位实数和虚数
	4	complex128：64 位实数和虚数
*/
/*
	其他数字类型
	以下列出了其他更多的数字类型：
	序号	类型和描述
	1	byte：类似 uint8
	2	rune：类似 int32
	3	uint：32 或 64 位
	4	int：与 uint 一样大小
	5	uintptr：无符号整型，用于存放一个指针
*/

func main() {

	fmt.Println("测试程序是否跑通")

}
