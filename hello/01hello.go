package main

import "fmt"

func init() {
	fmt.Println("this is 01hello.go # init")
}

func sayHello(name string) {
	fmt.Printf("hello! %s", name)
}
