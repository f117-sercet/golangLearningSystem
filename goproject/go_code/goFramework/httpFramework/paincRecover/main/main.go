package main

import (
	"golangLearningSystem/goproject/go_code/goFramework/httpFramework/paincRecover/gees"
	"net/http"
)

func main() {
	r := gees.Default()
	r.GET("/", func(c *gees.Context) {
		c.String(http.StatusOK, "Hello,World\n")
	})

	// 测试内存溢出
	r.GET("/panic", func(c *gees.Context) {
		names := []string{"hello,world"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")

}
