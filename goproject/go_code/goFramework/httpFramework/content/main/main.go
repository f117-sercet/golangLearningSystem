package main

import (
	"golangLearningSystem/goproject/go_code/goFramework/httpFramework/content/gees"
	"net/http"
)

func main() {

	r := gees.New()
	r.GET("/", func(c *gees.Content) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gees.Content) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gees.Content) {
		c.JSON(http.StatusOK, gees.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
