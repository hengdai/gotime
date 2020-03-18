package main

import (
	"fmt"
	"gotime/gotime"
)

func main()  {
	// 入参为时区，传入空字符串默认为 Asia/Shanghai
	gt := gotime.NewGoTime("")
	// 返回当前时间戳
	timestamps := gt.Timestamps()
	fmt.Println(timestamps)
}
