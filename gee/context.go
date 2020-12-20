package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H map
type H map[string]interface{}

// Context gee http context
type Context struct {
	Writer http.ResponseWriter
	*http.Request
	// request info
	StatusCode int

	Params map[string]string

	// middleware
	handlers []HandlerFunc
	index    int
}

func newContext(w http.ResponseWriter, req *http.Request, handlers ...HandlerFunc) *Context {
	return &Context{
		Writer:   w,
		Request:  req,
		handlers: handlers,
		index:    -1,
	}
}

func (c *Context) Next() {
	c.index++
	for ; c.index < len(c.handlers); c.index++ {
		c.handlers[c.index](c)
	}
}

// PostForm 获取表单
func (c *Context) PostForm(key string) string {
	return c.FormValue(key)
}

// Query 路由参数
func (c *Context) Query(key string) string {
	return c.URL.Query().Get(key)
}

// Param 获取Param参数
func (c *Context) Param(key string) string {
	return c.Params[key]
}

// Status http状态码写入
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 写入请求头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) Stringf(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 写入json
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 写入data
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 返回html
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
