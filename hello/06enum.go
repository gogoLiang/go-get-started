package main

import "fmt"

// TaskType 任务类型 枚举
// 枚举避免从0开始，除非零值与0有相同的意义
type TaskType int

const (
	// ExportWaybill 导出面单
	ExportWaybill TaskType = 1 + iota
	// ExportOrder 订单导出
	ExportOrder
	// ExportProduct 商品导出
	ExportProduct
)

func main() {

	fmt.Println(ExportWaybill)
	fmt.Println(ExportOrder)
	fmt.Println(ExportProduct)

}
