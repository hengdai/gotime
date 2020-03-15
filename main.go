package main

import (
	"fmt"
	"gotime/gotime"
)

func main()  {
	gt := gotime.NewGoTime("")
	timeStampsStr := gt.RTimestamps("2020-03-16 18:57:43")
	fmt.Println(timeStampsStr)
	fStr := gt.FTimestampsByRule(timeStampsStr, "y-m-d h:i:s")
	fmt.Println(fStr)
}
