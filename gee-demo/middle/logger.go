package middle

import (
	"log"
	"time"

	"github.com/NothingXiang/go-every-week/gee"
)

func Logger() gee.HandlerFunc {
	return func(c *gee.Context) {
		start := time.Now().UTC()

		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.RequestURI, time.Since(start))

	}
}
