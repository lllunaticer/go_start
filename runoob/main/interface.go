package main

/*
 Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
 接口可以让我们将不同的类型绑定到一组公共的方法上，从而实现多态和灵活的设计。
 Go 语言中的接口是隐式实现的，也就是说，如果一个类型实现了一个接口定义的所有方法，那么它就自动地实现了该接口。
因此，我们可以通过将接口作为参数来实现对不同类型的调用，从而实现多态。
*/
import (
	"fmt"
)

/*
Phone
我们定义了一个接口 Phone，接口里面有一个方法 call()。
然后我们在 main 函数里面定义了一个 Phone 类型变量，
并分别为之赋值为 NokiaPhone 和 IPhone。然后调用 call() 方法
*/
type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func testPhone() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}

// Shape
/*
我们定义了一个 Shape 接口，
它定义了一个方法 area()，该方法返回一个 float64 类型的面积值。
然后，我们定义了两个结构体 Rectangle 和 Triangle，它们分别实现了 Shape 接口的 area() 方法。
在 main() 函数中，我们首先定义了一个 Shape 类型的变量 s，
然后分别将 Rectangle 和 Triangle 类型的实例赋值给它，并通过 area() 方法计算它们的面积并打印出来，
*/
type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Triangle struct {
	bottom float64
	height float64
}

func (c Triangle) area() float64 {
	return 0.5 * c.bottom * c.height
}

func main() {
	var s Shape

	s = Rectangle{width: 10, height: 5}
	fmt.Printf("矩形面积: %f\n", s.area())

	s = Triangle{bottom: 3, height: 4}
	fmt.Printf("三角形形面积: %f\n", s.area())
}
