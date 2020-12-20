package main

import (
	"fmt"
	"net/http"

	"gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.Stringf(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.URL.Path)
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.Stringf(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.URL.Path)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	if err := r.Run(":8000"); err != nil {
		fmt.Println(err)
	}
}
