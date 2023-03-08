package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/utils"
)

func main() {
	list := arraylist.New()
	list.Add("a")
	list.Add("b", "c", "1")
	fmt.Println("new list:", list)
	list.Sort(utils.StringComparator)
	fmt.Println("after sort:", list)
	fmt.Println(list.Get(0))
	fmt.Println(list.Get(10))
	fmt.Println(list.Contains("b", "c"))
	l := list.Map(func(index int, value interface{}) interface{} {
		var v = value.(string)
		return v + v
	})
	fmt.Println("after map:", l)

	l2 := list.Select(func(index int, value interface{}) bool {
		v := value.(string)
		return v > "1"
	}).Map(func(index int, value interface{}) interface{} {
		var v = value.(string)
		return v + v
	})
	fmt.Println("after select map:", l2)

	values := l2.Values()
	l3 := arraylist.New(values)
	fmt.Println("after select map:", l3)
}
