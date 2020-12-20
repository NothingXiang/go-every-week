package gee

import (
	"log"
	"net/http"
)

// Router 自定义路由
type Router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']

func newRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc)}
}

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *Router) handle(c *Context) {
	key := c.Method + "-" + c.URL.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.URL.Path)
	}
}
