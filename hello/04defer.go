package main

import "fmt"

// defer 类似java的finally （返回前执行）
// defer 修饰的操作, 会在方法return前执行，单不印象return的值
// defer 通常用于执行资源释放, 尽量在申请资源后立即使用defer释放资源
// 使用defer增加可读性，并避免忘记释放资源
func main() {
	fmt.Println(execute())
}

func execute() int {
	add := func(a *int) {
		*a++
		fmt.Println("add done")
	}
	count := 0
	defer add(&count)
	add(&count)
	return count
}
