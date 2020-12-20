package gee

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup // support nesting
	engine      *Engine      // all groups share a Engine instance
}
