package main

import (
	"github.com/astaxie/beego/logs"
	"go-time/gotime"
)

func main()  {
	gt := gotime.NewGoTime("Asia/Shanghai")
	logs.Info(gt.FCorrByRule("y.m.d h.i.s"))
}
