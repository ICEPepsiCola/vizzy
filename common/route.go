package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouteContext struct {
	*gin.Context
}
type ResJSON struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func (c RouteContext) Success(data any) {
	c.JSON(http.StatusOK, ResJSON{
		Code: 0,
		Data: data,
		Msg:  "OK",
	})
}

func (c RouteContext) Fail(resJSON ResJSON) {
	c.JSON(http.StatusOK, resJSON)
}

func CustomizeHandler(c *gin.Context, fn func(cc *RouteContext)) {
	var cc = RouteContext{
		c,
	}
	fn(&cc)
}
