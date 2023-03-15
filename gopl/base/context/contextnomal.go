package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//context.TODO()
	var dealLock chan struct{} = nil
	background := context.Background()
	fmt.Printf("%d\r\n", &background)
	ctx, cancelFunc := context.WithCancel(background)
	//context.WithDeadline()
	fmt.Println(ctx.Deadline())
	ctx = context.WithValue(ctx, "user", "liangwenhui")
	for i := 0; i < 2; i++ {
		go func(ctx context.Context, id int) {
			for {
				select {
				case <-ctx.Done():
					fmt.Println(id, "正常结束", ctx.Value("user"))
					return
				default:
					fmt.Println(id, "运行中")
					time.Sleep(2 * time.Second)

				}
			}
		}(ctx, i)
	}
	dealLock <- struct{}{}
	fmt.Println("停止任务！")
	<-dealLock
	time.Sleep(4 * time.Second)
	fmt.Println("停止任务！")
	cancelFunc()
	fmt.Println(ctx.Err().Error())
	time.Sleep(2 * time.Second)
}
