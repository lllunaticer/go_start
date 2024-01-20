package main

import "fmt"

func array() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	var l, m, k int
	// 声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 输出数组元素 */
	for l = 0; l < 5; l++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出每个数组元素的值 */
	for m = 0; m < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}
}

func multiArray() {
	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	// Step 3: 显示两行数据
	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// Step 4: 访问第一个元素
	fmt.Println("第一个元素为：")
	fmt.Println(values[0][0])

	// 创建二维数组
	sites := [2][2]string{}

	// 向二维数组添加元素
	sites[0][0] = "Google"
	sites[0][1] = "Runoob"
	sites[1][0] = "Taobao"
	sites[1][1] = "Weibo"

	// 显示结果
	fmt.Println(sites)

	/* 数组 - 5 行 2 列*/
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}

	// 创建空的二维数组
	animals := [][]string{}

	// 创建三一维数组，各数组长度不同
	row4 := []string{"fish", "shark", "eel"}
	row5 := []string{"bird"}
	row3 := []string{"lizard", "salamander"}

	// 使用 append() 函数将一维数组添加到二维数组中
	animals = append(animals, row4)
	animals = append(animals, row5)
	animals = append(animals, row3)

	// 循环输出
	for i := range animals {
		fmt.Printf("Row: %v\n", i)
		fmt.Println(animals[i])
	}
}

// 数组作为函数参数
// 让我们看下以下实例，实例中函数接收整型数组参数，另一个参数指定了数组元素的个数，并返回平均值
func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum / size)
	return avg
}

func testGetAverage2() {
	/* 数组长度为 5 */
	var balance = [5]int{1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	avg = getAverage2(balance, 5)

	/* 输出返回的平均值 */
	fmt.Printf("平均值为: %f ", avg)
}

func getAverage2(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum) / float32(size)
	return avg
}

// 浮点数计算输出有一定的偏差，你也可以转整型来设置精度
func precise() {
	a := 1.69
	b := 1.7
	c := a * b     // 结果应该是2.873
	fmt.Println(c) // 输出的是2.8729999999999998

	d := 1690                         // 表示1.69
	e := 1700                         // 表示1.70
	f := d * e                        // 结果应该是2873000表示 2.873
	fmt.Println(f)                    // 内部编码
	fmt.Println(float64(f) / 1000000) // 显示
}

// 如果你想要在函数内修改原始数组，可以通过传递数组的指针来实现。
// 以下实例演示如何向函数传递数组，函数接受一个数组和数组的指针作为参数：
// 函数接受一个数组作为参数
func modifyArray(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
}

// 函数接受一个数组的指针作为参数
func modifyArrayWithPointer(arr *[5]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] * 2
	}
}

func testModifyArray() {
	// 创建一个包含5个元素的整数数组
	myArray := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Original Array:", myArray)

	// 传递数组给函数，但不会修改原始数组的值
	modifyArray(myArray)
	fmt.Println("Array after modifyArray:", myArray)

	// 传递数组的指针给函数，可以修改原始数组的值
	modifyArrayWithPointer(&myArray)
	fmt.Println("Array after modifyArrayWithPointer:", myArray)
}
