package main

import (
	"fmt"
	"gotime/gotime"
)

func main()  {
	gt := gotime.NewGoTime("Asia/Shanghai")
	gotime.Sleep(3)
	fmt.Println(gt.FCurrDefault())
}
