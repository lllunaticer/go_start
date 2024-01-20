package main

import "fmt"

func partialVisible() {
	/* 声明局部变量 */
	var a, b, c int

	/* 初始化参数 */
	a = 10
	b = 20
	c = a + b

	fmt.Printf("结果： a = %d, b = %d and c = %d\n", a, b, c)
}

/* 声明全局变量 */
var g int

func globalVisible() {
	/* 声明局部变量 */
	var a, b int

	/* 初始化参数 */
	a = 10
	b = 20
	g = a + b

	fmt.Printf("结果： a = %d, b = %d and g = %d\n", a, b, g)
}
