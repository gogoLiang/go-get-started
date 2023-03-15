package main

import (
	"fmt"
	"time"
)

func main() {
	closeSign := make(chan struct{})
	for i := 0; i < 2; i++ {
		go func() {
			select {
			case <-closeSign:
				fmt.Println("closed!")
			}
		}()
	}
	//close(closeSign) 将管道关闭, 从closeSign获取值都会直接获取管道类型的零值
	close(closeSign)
	time.Sleep(1 * time.Second)
}
