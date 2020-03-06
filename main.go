package main

import (
	"github.com/astaxie/beego/logs"
	"gotime/gotime"
)

func main()  {
	//gt := gotime.NewGoTime("Asia/Shanghai")
	//logs.Info(gt.FCorrByRule("y.m.d h.i.s"))
	logs.Info(gotime.FTimestampsByRule(1573304084, "y-m-d h:i:s"))
}
