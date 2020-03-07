package main

import (
	"github.com/astaxie/beego/logs"
	"gotime/gotime"
)

func main()  {
	gt := gotime.NewGoTime("Asia/Shanghai")
	logs.Info(gt.FCurrDefault())
	gotime.Sleep(3)
	logs.Info(gt.FCurrDefault())
}
