package main

import (
	"errors"
	"fmt"
)

// 如果调用方需要知道错误类型，并处理，必须声明顶级错误变量
// 或自定义类型来支持errors.IS 或 errors.As
// 错误的消息是否是静态的字符串，还是需要动态的字符串。
// 静态字符串错误通过 errors.New, 动态使用 fmt.Errorf
// panic(err.Error()) painc会退出程序
func main() {

	err := Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = OpenFile("one_file")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = Create()
	if errors.Is(err, CreateError) {
		fmt.Println(err)
	}
	err = CreateFile("tow_file")
	var tarErr *CreateFileError
	if errors.As(err, &tarErr) {
		fmt.Println(err)
	}
}

func Open() error {
	return errors.New("system error")
}

func OpenFile(fileName string) error {
	return fmt.Errorf("%s not found", fileName)
}

var CreateError = errors.New("system create error")

func Create() error {
	return CreateError
}

type CreateFileError struct {
	message string
}

// error接口 Error方法, 记得要使用值接收器，才是完全实现接口
func (e CreateFileError) Error() string {
	return fmt.Sprintf("%s create error", e.message)
}

func CreateFile(fn string) error {
	return &CreateFileError{message: fn}
}
