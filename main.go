package main

import (
	"fmt"
	"gotime/gotime"
	"time"
)

func main()  {
	gt := gotime.NewGoTime("")
	t := time.Now()
	res := gt.Shift(t, "5years", "y-m-d h:i:s")
	fmt.Println(res)
}
