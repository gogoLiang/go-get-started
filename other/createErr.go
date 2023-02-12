package other

import (
	"errors"
	"fmt"
)

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

type createFileError struct {
	message string
}

// error接口 Error方法, 记得要使用值接收器，才是完全实现接口
func (e createFileError) Error() string {
	return fmt.Sprintf("%s create error", e.message)
}

func New(fn string) createFileError {
	return createFileError{fn}
}

func CreateFile(fn string) error {
	return createFileError{fn}
}
