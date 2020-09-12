package main

import (
	"fmt"
	"time"

	"github.com/NothingXiang/go-every-week/jsonops"
)

func main() {
	start := time.Now()
	fmt.Printf("start:%v\n", start)

	jsonops.GetValueDemo()

	end := time.Now()
	fmt.Printf("end:%v,cost:%v\n", end, end.Sub(start))
}
