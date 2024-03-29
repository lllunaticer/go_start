package main

import "fmt"

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

func constant() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"
	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)

}

/*
iota，特殊常量，可以认为是一个可以被编译器修改的常量。
iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
iota 可以被用作枚举值：
*/
const (
	a = iota
	b = iota
	c = iota
)

// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
const (
	d = iota
	e
	f
)

func iotaVar() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	const (
		l = 1 << iota
		m = 3 << iota
		n
		o
	)
	fmt.Println("l=", l)
	fmt.Println("m=", m)
	fmt.Println("n=", n)
	fmt.Println("o=", o)
}
