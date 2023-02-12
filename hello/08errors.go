package main

import (
	"errors"
	"fmt"
	"go-get-started/other"
)

// 如果调用方需要知道错误类型，并处理，必须声明顶级错误变量
// 或自定义类型来支持errors.IS 或 errors.As
// 错误的消息是否是静态的字符串，还是需要动态的字符串。
// 静态字符串错误通过 errors.New, 动态使用 fmt.Errorf
// panic(err.Error()) painc会退出程序
func main() {

	err := other.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = other.OpenFile("one_file")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = other.Create()
	if errors.Is(err, other.CreateError) {
		fmt.Println(err)
	}
	cerr := other.CreateFile("tow_file")
	fileError := other.New("")
	fmt.Println(cerr == nil)
	if errors.As(cerr, &fileError) {
		fmt.Println(cerr)
	}

	var t1 T = "123"
	var t2 T = "456"
	t1.hello(t2)

}

type T string

func (t *T) hello(t2 T) {
	fmt.Println("hello", t, t2)
}
