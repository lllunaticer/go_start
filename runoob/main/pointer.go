package main

import "fmt"

func pointer() {
	var a int = 10
	fmt.Printf("变量的地址: %x\n", &a)

	var ip *int /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

/*
a 变量的地址是: 20818a220
ip 变量储存的指针地址: 20818a220
*ip 变量的值: 20
*/

func nullPointer() {
	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)

	/* ptr 不是空指针 */
	if ptr != nil {
		fmt.Println()
	}
	/* ptr 是空指针 */
	if ptr == nil {
		fmt.Println()
	}
}

const MAX int = 3

func pointerArray() {
	a := []int{10, 100, 200}
	var i int

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}

	b := []int{10, 100, 200}
	var j int
	var ptr [MAX]*int

	for j = 0; j < MAX; j++ {
		ptr[i] = &b[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("b[%d] = %d\n", j, *ptr[j])
	}
}

func pointerPointer() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}

// 输出
//变量 a = 3000
//指针变量 *ptr = 3000
//指向指针的指针变量 **pptr = 3000
