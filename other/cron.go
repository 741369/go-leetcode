/**********************************************
** @Des:
** @Author: 1@lg1024.com
** @Last Modified time: 2020/3/21 下午11:29
***********************************************/

package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	//"github.com/robfig/cron"
)

func main() {
	runCron()
	select {}
}

func runCron() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0, 0, 12, *, *, *", fmtTest) // 定时推送消息

	c.AddFunc("*/1, *, *, *, *, *", func() {
		fmt.Println("===", time.Now().Format(time.RFC3339))
	}) // 定时推送消息
	c.Start()
}

//5分钟执行
func run1130EveryWeekend() {
	// 秒，分，时，日，月，星期
	spec := "*/5, *, *, *, *, *" // 每周五11：30
	c := cron.New(cron.WithSeconds())
	c.AddFunc(spec, fmtTest) // 定时推送消息
	c.Start()
}

func run1Second() {
	spec := "*/1, *, *, *, *, *" // 每秒
	c := cron.New(cron.WithSeconds())
	c.AddFunc(spec, func() {
		fmt.Println("===", time.Now().Format(time.RFC3339))
	}) // 定时推送消息
	c.Start()
}

func fmtTest() {
	fmt.Println("==+++++==", time.Now().Format(time.RFC3339))
}
