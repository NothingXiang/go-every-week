package gee

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	routerKey = "%v-%v"
)

// Router 自定义路由
// 概括一下，router的功能应该有2大类： 1是插入路由，2是匹配路由
type Router struct {
	// 每种请求方式的根节点
	// roots key eg, roots['GET'] roots['POST']
	roots map[string]*node

	// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// parsePattern
// Only one * is allowed ,after '*' will be ignore
func parsePattern(pattern string) (parts []string) {
	paths := strings.Split(pattern, "/")
	parts = make([]string, 0)

	for _, path := range paths {
		if path != "" {
			parts = append(parts, path)
			if path[0] == '*' {
				break
			}
		}
	}
	return
}

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	/*log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler*/

	parts := parsePattern(pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[fmt.Sprintf(routerKey, method, pattern)] = handler
}

// getRoute 根据method和url匹配对应的处理节点
// param: method	e.g: GET, POST
// param: pattern	URL.Path
// return:
//	node: 匹配到的节点，用于查找对应的handler
//	map: params参数
func (r *Router) getRoute(method string, pattern string) (node *node, params map[string]string) {
	// 实际访问路径
	extraParts := parsePattern(pattern)

	// 判断是否有该方法的路由
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	// 查找路由实际匹配的节点
	node = root.search(extraParts, 0)
	if node == nil {
		return nil, nil
	}

	// 获取params参数
	params = make(map[string]string)
	parts := parsePattern(node.pattern)
	for idx, part := range parts {
		if part[0] == ':' {
			params[part[1:]] = extraParts[idx]
			continue
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(extraParts[idx:], "/")
			break
		}
	}

	return node, params
}

func (r *Router) handle(c *Context) {

	// match node
	node, paramas := r.getRoute(c.Method, c.URL.Path)
	if node == nil {
		c.handlers = append(c.handlers, func(c *Context) {
			c.Stringf(http.StatusNotFound, "404 NOT FOUND: %s\n", c.URL.Path)
		})
	} else {
		// get handler and params
		c.Params = paramas
		key := fmt.Sprintf(routerKey, c.Method, node.pattern)
		c.handlers = append(c.handlers, r.handlers[key])
	}

	// handle request 保证这个入口是第一次调用c.Next()
	c.Next()
}
