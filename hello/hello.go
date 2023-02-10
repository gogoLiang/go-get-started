package main

import "fmt"

func init() {
	fmt.Println("this is hello.go # init")
}

func sayHello(name string) {
	fmt.Printf("hello! %s", name)
}
