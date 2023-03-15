package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func MockPdf(_ context.Context, c *app.RequestContext) {

	c.File("F:\\谷歌下载\\Bluetooth 5.0 Adapter Installation V1\\Bluetooth 5.0 Adapter Installation\\安装指南.pdf")

}
