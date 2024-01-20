package main

import (
	"fmt"
	"math"
)

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

// 值传递	值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
/* 定义相互交换值的函数 */
func swapValue(x, y int) int {
	var temp int
	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/
	return temp
}

/*
交换前 a 的值为 : 100
交换前 b 的值为 : 200
交换后 a 的值 : 100
交换后 b 的值 : 200
*/
func swapValueTest() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	/* 通过调用函数来交换值 */
	swapValue(a, b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

// 引用传递	引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
func swapRef(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

func swapRefTest() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前，a 的值 : %d\n", a)
	fmt.Printf("交换前，b 的值 : %d\n", b)

	/* 调用 swap() 函数
	 * &a 指向 a 指针，a 变量的地址
	 * &b 指向 b 指针，b 变量的地址
	 */
	swapRef(&a, &b)

	fmt.Printf("交换后，a 的值 : %d\n", a)
	fmt.Printf("交换后，b 的值 : %d\n", b)
}

/*
Go 语言可以很灵活的创建函数，并作为另外一个函数的实参。以下实例中我们在定义的函数中初始化一个变量，该函数仅仅是为了使用内置函数 math.sqrt()，实例为：
*/
func funAsParams() {
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9))
}

/*
Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
匿名函数是一种没有函数名的函数，通常用于在函数内部定义函数，或者作为函数参数进行传递。
以下实例中，我们创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量，代码如下
*/
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/*
以下代码执行结果为：
1
2
3
1
2
*/
func getSequenceTest() {
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
}

func anonymousFunc() {
	// 定义一个匿名函数并将其赋值给变量add
	add := func(a, b int) int {
		return a + b
	}

	// 调用匿名函数
	result := add(3, 5)
	fmt.Println("3 + 5 =", result)

	// 在函数内部使用匿名函数
	multiply := func(x, y int) int {
		return x * y
	}

	product := multiply(4, 6)
	fmt.Println("4 * 6 =", product)

	// 将匿名函数作为参数传递给其他函数
	calculate := func(operation func(int, int) int, x, y int) int {
		return operation(x, y)
	}

	sum := calculate(add, 2, 8)
	fmt.Println("2 + 8 =", sum)

	// 也可以直接在函数调用中定义匿名函数
	difference := calculate(func(a, b int) int {
		return a - b
	}, 10, 4)
	fmt.Println("10 - 4 =", difference)
}

/*
以上代码执行结果为：
3 + 5 = 8
4 * 6 = 24
2 + 8 = 10
10 - 4 = 6
*/

/*Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，
接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
所有给定类型的方法属于该类型的方法集*/

/* 定义结构体 */
type Circle struct {
	radius float64
}

// 该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func testGetArea() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}
