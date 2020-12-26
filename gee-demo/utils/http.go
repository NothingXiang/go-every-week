package utils

import (
	"net/http"

	"github.com/NothingXiang/go-every-week/gee"
)

// SucResp
func SucResp(c *gee.Context, data interface{}) {
	c.JSON(http.StatusOK,
		gee.H{
			"code": 0,
			"msg":  "",
			"data": data,
		},
	)
}

// ErrResp
func ErrResp(c *gee.Context, code int, err error) {
	c.JSON(code, gee.H{
		"code": code,
		"msg":  err.Error(),
	})
}
