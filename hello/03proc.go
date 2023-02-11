package main

import (
	"log"
	"os"
	"sync"
	"time"
)

func init() {
	log.SetOutput(os.Stdout)
}

/*
*
模拟一个程序执行流程
1. 建立一个协程A获取数据
2. 建立另一个协程B等待数据
3. B处理完数据后,主线程打印数据
*/
type DataResult struct {
	code int
	data interface{}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	c := make(chan *DataResult, 1)

	go func() {
		log.Println("waiting for search api")
		wg.Wait()
		log.Println("a done! ")

	}()

	go func() {
		s := search()
		c <- &DataResult{code: 200, data: s}
		wg.Done()
	}()

	dr := <-c
	log.Printf("%v", *dr)

}

func search() string {
	time.Sleep(time.Second * 5)
	return `{"key":"gogoliang"}`
}
