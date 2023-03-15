package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

type createOrder struct {
	OrderNo string `json:"order_no"`
}

func Mock3pl(_ context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, createOrder{"1234568LL"})
}
