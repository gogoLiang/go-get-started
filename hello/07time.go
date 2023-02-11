package main

import (
	"fmt"
	"time"
)

// 标准库处理时间
func main() {

	now := time.Now()
	fmt.Println(now)

	start := time.Date(2023, 2, 10, 0, 0, 0, 0, time.Local)
	stop := time.Date(2023, 2, 14, 0, 0, 0, 0, time.Local)
	now.Add(time.Hour * 24)
	fmt.Println(now)
	fmt.Println("start before now", start.Before(now), "stop after now", stop.After(now))

}
