package main

import "fmt"

// 切片底层使用数据储存数据, 切片记录了从数组中
// 某一个元素开始,到另一个元素结束（容量）的可操作范围
// 切片、map，或者说指针类型，要考虑是否能被用户修改, 不期望被
// 修改的情况下，请复制后传递数据，避免暴露
func main() {

	// 定义一个数组, 容量为4, 字面量赋值了3个元素，剩余会初始化类型的零值
	// 只有定义了数组长度，才是数组 [] 则为切片
	arr1 := [4]int{1, 2, 3}
	// 1 2 3 0
	fmt.Println("arr1:", arr1)
	//定义切片
	// make([]type, len, cap)
	// []type{}
	// 通过数组 arr[startIndex:endIndex] 不填默认首尾
	slice1 := make([]int, 10)
	slice2 := []int{1, 2}
	slice3 := arr1[1:]
	fmt.Println("s1:", slice1)
	fmt.Println("s2:", slice2)
	fmt.Println("s3:", slice3)

	//通过修改数组 、 切片，是否会相互影响
	//数组arr1修改下标1 - > 10
	arr1[1] = 10
	fmt.Println("s3[arr1.1 -> 10]:", slice3)
	slice3[0] = 20
	fmt.Println("arr[s3.0 -> 20]:", arr1)
	//当s3 扩容，是将arr1复制到新数组，还是记录2个数组
	slice4 := append(slice3, 5)
	fmt.Println("s4:", slice4)
	slice4[2] = 4
	slice3[0] = 2
	fmt.Println("s3[s4.2->4]", slice3)
	fmt.Println("s4[s4.2->4]", slice4)
	fmt.Println("arr[s4.2->4]", arr1)
	// s3 s4 已经不共享数据了

	//当切片通过数组初始化，则容量为len(arr) - startIndex
	slice5 := arr1[2:3]
	fmt.Println("s5", slice5, cap(slice5))
	var _ = append(slice5, 100)
	fmt.Println("arr", arr1)
	//切片零值 nil []
	var nilSlice []int
	fmt.Println(nilSlice, nilSlice == nil)

}
