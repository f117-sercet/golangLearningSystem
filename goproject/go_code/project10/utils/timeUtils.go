package utils

import (
	"fmt"
	"time"
)

func Test() {

	// 获取当前时间
	now := time.Now()
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 格式化日期格式
	// 数字固定，组合随意
	now.Format("2006-01-02 15:04:05")
	unix := now.UnixNano()
	fmt.Println(unix)
}
