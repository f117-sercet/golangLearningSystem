package middleWare

import (
	"log"
	"time"
)

func Logger() HandlerFunc {

	result := func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
	return result
}
