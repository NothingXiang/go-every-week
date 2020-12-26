package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/NothingXiang/go-every-week/gee"
	cache "github.com/NothingXiang/go-every-week/gee-cache"
	"github.com/NothingXiang/go-every-week/gee-demo/middle"
	"github.com/NothingXiang/go-every-week/gee-demo/utils"
)

const (
	defaultBasePath = "/_geecache"
)

var (
	_ = cache.NewGroup("score", 2<<10, cache.GetterFunc(LoadData))
)

func main() {
	defer func() {
		log.Println("recover:", recover())
	}()

	engine := gee.New()

	engine.Use(middle.Logger())
	cacheGroup := engine.Group(defaultBasePath)
	{
		cacheGroup.GET("/", Get)
		// todo: set cache
		cacheGroup.POST("/", Set)
	}

	var port string
	flag.StringVar(&port, "port", "8040", "server port")
	flag.Parse()

	if err := engine.Run(":" + port); err != nil {
		log.Fatalln(err)
	}
}

func Set(c *gee.Context) {
	// do nothing
}

func Get(c *gee.Context) {
	groupName := c.Query("group")
	key := c.Query("key")

	group := cache.GetGroup(groupName)

	if group == nil {
		utils.ErrResp(c, http.StatusNotFound, fmt.Errorf("unknown group name:%s", groupName))
		return
	}

	view, err := group.Get(key)
	if err != nil {
		//gee_demo
		utils.ErrResp(c, http.StatusNotFound, err)
		return
	}

	c.SetHeader("Content-Type", "application/octet-stream")
	utils.SucResp(c, view.String())
}

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func LoadData(key string) ([]byte, error) {
	log.Println("[SlowDB] search key", key)
	if v, ok := db[key]; ok {
		return []byte(v), nil
	}
	return nil, fmt.Errorf("%s not exist", key)
}
