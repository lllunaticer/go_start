package main

import "fmt"

func operator() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)

	if a == b {
		fmt.Printf("第一行 - a 等于 b\n")
	} else {
		fmt.Printf("第一行 - a 不等于 b\n")
	}
	if a < b {
		fmt.Printf("第二行 - a 小于 b\n")
	} else {
		fmt.Printf("第二行 - a 不小于 b\n")
	}

	if a > b {
		fmt.Printf("第三行 - a 大于 b\n")
	} else {
		fmt.Printf("第三行 - a 不大于 b\n")
	}
	/* Lets change value of a and b */
	a = 5
	b = 20
	if a <= b {
		fmt.Printf("第四行 - a 小于等于 b\n")
	}
	if b >= a {
		fmt.Printf("第五行 - b 大于等于 a\n")
	}

	var e bool = true
	var d bool = false
	if e && d {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if e || d {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 a 和 b 的值 */
	e = false
	d = true
	if e && d {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(e && d) {
		fmt.Printf("第四行 - 条件为 true\n")
	}

	var f uint = 60 /* 60 = 0011 1100 */
	var g uint = 13 /* 13 = 0000 1101 */
	var h uint = 0

	h = f & g /* 12 = 0000 1100 */
	fmt.Printf("第一行 - h 的值为 %d\n", h)

	h = f | g /* 61 = 0011 1101 */
	fmt.Printf("第二行 - h 的值为 %d\n", h)

	h = f ^ g /* 49 = 0011 0001 */
	fmt.Printf("第三行 - h 的值为 %d\n", h)

	h = f << 2 /* 240 = 1111 0000 */
	fmt.Printf("第四行 - h 的值为 %d\n", h)

	h = f >> 2 /* 15 = 0000 1111 */
	fmt.Printf("第五行 - h 的值为 %d\n", h)

	var i int = 4
	var j int32
	var k float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - i 变量类型为 = %T\n", i)
	fmt.Printf("第 2 行 - j 变量类型为 = %T\n", j)
	fmt.Printf("第 3 行 - k 变量类型为 = %T\n", k)

	/*  & 和 * 运算符实例 */
	ptr = &i /* 'ptr' 包含了 'i' 变量的地址 */
	fmt.Printf("i 的值为  %d\n", i)
	fmt.Printf("*ptr 为 %d\n", *ptr)

}
