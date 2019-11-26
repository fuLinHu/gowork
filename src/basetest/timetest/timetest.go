package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Second())

	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println("------------------------------------")

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	unix := time.Unix(1574748742, 0)
	fmt.Println(unix)
	fmt.Println("------------------------------------")
	tick := time.Tick(time.Second * 2)
	formatDemo()
	for i := range tick {
		fmt.Println(i)
	}

}

func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println("---------------日期格式化------------------")
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println("---------------日期格式化------------------")
}
