package gee

import (
	"fmt"
	"net/http"
	"strings"
)

// HandlerFunc http处理器
type HandlerFunc func(*Context)

// Engine http engine
type Engine struct {
	router *Router
	*RouterGroup
	groups []*RouterGroup // store all groups
}

// New is the constructor of gee.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Run defines the method to start a http server
func (e *Engine) Run(addr string) (err error) {
	fmt.Println("start run")

	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(w, req, middlewares...)
	e.router.handle(c)
}
