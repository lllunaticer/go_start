package main

import (
	"fmt"
	"strconv"
)

func typeConvert() {

	var sum int = 17
	var count int = 5
	var mean float32
	//整型转化为浮点型
	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)

	//字符串转换成整数
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
	}

	// 整数转字符串
	num1 := 123
	str1 := strconv.Itoa(num1)
	fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num1, str1)

	//字符串转浮点数
	str2 := "3.14"
	num2, err := strconv.ParseFloat(str2, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str2, num2)
	}

	//浮点数转换为字符串：
	num3 := 3.14
	str3 := strconv.FormatFloat(num3, 'f', 2, 64)
	fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num3, str3)

	//接口类型转换有两种情况：类型断言和类型转换。
	//类型断言用于将接口类型转换为指定类型，其语法为
	//	value.(type)
	//或者
	//value.(T)

	/*以下，我们定义了一个接口类型变量 i，并将它赋值为字符串 "Hello, World"。
	然后，我们使用类型断言将 i 转换为字符串类型，并将转换后的值赋值给变量 str。
	最后，我们使用 ok 变量检查类型转换是否成功，
	如果成功，我们打印转换后的字符串；否则，我们打印转换失败的消息。*/
	var i interface{} = "Hello, World"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
}

/*
	类型转换用于将一个接口类型的值转换为另一个接口类型，其语法为
	T(value)
	T 是目标接口类型，value 是要转换的值。
	在类型转换中，我们必须保证要转换的值和目标接口类型之间是兼容的，否则编译器会报错。

以下实例中，我们定义了一个 Writer 接口和一个实现了该接口的结构体 StringWriter。
然后，我们将 StringWriter 类型的指针赋值给 Writer 接口类型的变量 w。接着，
我们使用类型转换将 w 转换为 StringWriter 类型，并将转换后的值赋值给变量 sw。
最后，我们使用 sw 访问 StringWriter 结构体中的字段 str，并打印出它的值。
*/

type Writer interface {
	Write([]byte) (int, error)
}

type StringWriter struct {
	str string
}

func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}

func testStringWriter() {
	var w Writer = &StringWriter{}
	sw := w.(*StringWriter)
	sw.str = "Hello, World"
	fmt.Println(sw.str)
}
