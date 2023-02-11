package main

import "fmt"

/**
go 函数参数传递是值传递
- 如果传递的参数是值类型,则在方法调用时，将原值得数据复制一份，对方法内的数据修改
不会影响原数据
- 在使用数组等可能较大的数据时,值传递会消耗大量资源
- 尽量使用指针作为入参
*/

type I interface {
	//I 接口有一个ifunc方法
	ifunc()
}

type A struct {
}

func (a A) ifunc() {
	fmt.Println("a.ifunc")
}

type B struct {
}

func (bp *B) ifunc() {
	fmt.Println("bp.ifunc")
}

func main() {
	var _ I = A{}
	var _ I = &A{}
	var _ I = &B{}
	//var _ I = B{} 编译错误

	A{}.ifunc()
	ap := &A{}
	ap.ifunc()
	// B{}.ifunc()编译失败
	(&B{}).ifunc()

	/*
		A B 两个结构体都实现了I接口，为什么B{} 无法赋值给 I类型的变量呢
		- 方法集
		类型的值 - 值方法集
		类型的指针 - 指针方法集
		func (a A) ifunc() {} 其中 a A 表示 A类型的值作为ifunc方法的接受者
		值作为接受者的方法，类型的值和指针都能调用方法，因为指针方法集包含值方法集
		使用指针是可以调用值接收器的方法的。因此 A实现的方法,值和指针都能匹配，所以A
		与I完全匹配
		而B定义了指针作为接收器，那只要指针才能调用这个方法，值不能调用。则只有&B才与I匹配
		指针类型可以修改底层数据, 并且复制消耗的性能较低
	*/
}
