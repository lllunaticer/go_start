package main

import "fmt"

var x, y int

var ( // 因式分解式关键字写法，用于声明全局变量
	xx int
	yy bool
)

func Variant() {
	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// 声明一个变量并初始化
	var d = "RUNOOB"
	fmt.Println(d)

	// 没有初始化就为零值
	var e int
	fmt.Println(e)

	// bool 零值为 false
	var f bool
	fmt.Println(f)

	var i int
	var g float64
	var h bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, g, h, s)

	//声明并赋值
	str := "string"
	fmt.Println(str)

	//等同于
	var str2 = "string"
	fmt.Println(str2)

	// 并行赋值
	//多变量可以在同一行进行赋值，如：
	var a1, a2 int
	var a3 string
	a1, a2, a3 = 1, 2, "string"
	fmt.Println(a1, a2, a3)

	// 空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃
	// _ 实际上是一个只写变量，你不能得到它的值
	_, b = 5, 7

}
