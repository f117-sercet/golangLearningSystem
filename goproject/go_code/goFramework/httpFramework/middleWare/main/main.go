package main

import (
	"golangLearningSystem/goproject/go_code/goFramework/httpFramework/middleWare"
	"log"
	"net/http"
	"time"
)

func onlyForv2() middleWare.HandlerFunc {

	return func(c *middleWare.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {

	r := middleWare.New()
	r.Use(middleWare.Logger())
	r.GET("/", func(context *middleWare.Context) {
		context.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForv2())
	{
		v2.GET("/hello/:name", func(c *middleWare.Context) {

			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}
	r.Run(":9999")

}
