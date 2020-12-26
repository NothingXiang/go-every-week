module github.com/NothingXiang/go-every-week/cache-server

replace (
	github.com/NothingXiang/go-every-week/gee => ../gee
	github.com/NothingXiang/go-every-week/gee-cache => ../gee-cache
	github.com/NothingXiang/go-every-week/gee-demo => ../gee-demo
)

go 1.15

require (
	github.com/NothingXiang/go-every-week/gee v0.0.0-20201224174056-7f4f62671d26
	github.com/NothingXiang/go-every-week/gee-cache v0.0.0-00010101000000-000000000000
	github.com/NothingXiang/go-every-week/gee-demo v0.0.0-00010101000000-000000000000
)
