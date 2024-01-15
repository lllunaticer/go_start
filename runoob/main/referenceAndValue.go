package main

import "fmt"

func ReferenceAndValue() {
	//	所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值
	//	当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝
	i := 1
	j := i

	fmt.Println(i, j)

	//	你可以通过 &i 来获取变量 i 的内存地址，例如：0xf840000040（每次的地址都可能不一样）。
	k := &i
	m := &j
	fmt.Println("the add of i is ", k, " the add of j is ", m)
}
