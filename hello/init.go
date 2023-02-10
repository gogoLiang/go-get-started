package main

import (
	"log"
	"os"
)

// init 每个init方法都会在main之前执行
func init() {
	//日志输出到std out
	log.SetOutput(os.Stdout)
	log.Println(`this is init.go # init func`)
}

func main() {
	log.Println("application starter")
	sayHello("admin")
}
