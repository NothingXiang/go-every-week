package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/NothingXiang/go-every-week/gee"
	"github.com/NothingXiang/go-every-week/gee-demo/middle"
)

func main() {
	r := gee.New()
	r.Use(middle.Logger())

	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=geektutu
			c.Stringf(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.URL.Path)
		})
	}
	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.Stringf(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.URL.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	if err := r.Run("0.0.0.0:8000"); err != nil {
		fmt.Println(err)
	}
}

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.RequestURI, time.Since(t).Milliseconds())
	}
}
