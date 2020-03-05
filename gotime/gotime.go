package gotime

import (
	"strconv"
	"strings"
	"time"
	"unicode"
)

type GoTime struct {
	Location *time.Location `json:"location"`
}

type SingleFormat struct {
	Len int `json:"len"`
	Suffix string `json:"suffix"`
	Index int `json:"index"`
}

var monthIntMap = map[string]string{
	"January": "01",
	"February": "02",
	"March": "03",
	"April": "04",
	"May": "05",
	"June": "06",
	"July": "07",
	"August": "08",
	"September": "09",
	"October": "10",
	"November": "11",
	"December": "12",
}

// 时区默认是Asia/Shanghai
func NewGoTime(v interface{}) *GoTime  {
	var gt GoTime
	if v, ok := v.(string); ok {
		gt.Location, _ = time.LoadLocation(v)
	} else {
		panic("please use string type for NewGoTime func")
	}
	return  &gt
}

// 返回秒级别时间戳
func (t *GoTime) Timestamps() int64 {
	timestamps := time.Now().In(t.Location)
	return timestamps.Unix()
}

// 返回纳秒时间
func (t *GoTime) Nanosecond() int {
	timestamps := time.Now().In(t.Location)
	return timestamps.Nanosecond()
}

// 返回秒加纳秒的时间戳
func (t *GoTime) TimestampsWithNano() string {
	timestamps := t.Timestamps()
	nanosecond := t.Nanosecond()
	timestampsWithNano := strconv.FormatInt(timestamps, 10) + "." + strconv.Itoa(nanosecond)
	return timestampsWithNano
}

// 格式化当前时间，默认的format格式是YYYY-MM-DD HH:MM:SS
func (t *GoTime) FCurrDefault() string {
	return time.Now().In(t.Location).Format("2006-01-02 15:04:05")
}

// 按照给定的格式化当前时间，例如：y-m-d h:i:s
func (t *GoTime) FCorrByRule(rule string) string {
	var year SingleFormat
	var month SingleFormat
	var day SingleFormat
	var hour SingleFormat
	var minute SingleFormat
	var second SingleFormat
	ruleArr := strings.Split(rule, "")
	timeNow := time.Now()

	for key, value := range rule {
		if unicode.IsLetter(value) {
			if strings.ContainsRune("yY", value) {
				year.Len += 1
				year.Index = key
			} else if strings.ContainsRune("mM", value) {
				month.Len += 1
				month.Index = key
			} else if strings.ContainsRune("dD", value) {
				day.Len += 1
				day.Index = key
			} else if strings.ContainsRune("hH", value) {
				hour.Len += 1
				hour.Index = key
			} else if strings.ContainsRune("iI", value) {
				minute.Len += 1
				minute.Index = key
			} else if strings.ContainsRune("sS", value) {
				second.Len += 1
				second.Index = key
			}
		} else {
			if key - year.Index == 1 {
				year.Suffix = ruleArr[key]
			} else if  key - month.Index == 1 {
				month.Suffix = ruleArr[key]
			} else if  key - day.Index == 1 {
				day.Suffix = ruleArr[key]
			} else if  key - hour.Index == 1 {
				hour.Suffix = ruleArr[key]
			} else if  key - minute.Index == 1 {
				minute.Suffix = ruleArr[key]
			} else if  key - second.Index == 1 {
				second.Suffix = ruleArr[key]
			}
		}
	}

	ret := ""
	yearStr := strconv.Itoa(timeNow.Year())
	if year.Len > 0 {
		ret += yearStr + year.Suffix
	}

	monthStr := monthIntMap[timeNow.Month().String()]
	if month.Len > 0 {
		ret += monthStr + month.Suffix
	}

	dayStr := strconv.Itoa(timeNow.Day())
	if len(dayStr) == 1 {
		dayStr = "0" + dayStr
	}
	if day.Len > 0 {
		ret += dayStr + day.Suffix
	}

	hourStr := strconv.Itoa(timeNow.Hour())
	if len(hourStr) == 1 {
		hourStr = "0" + hourStr
	}
	if hour.Len > 0 {
		ret += hourStr + hour.Suffix
	}

	minuteStr := strconv.Itoa(timeNow.Minute())
	if len(minuteStr) == 1 {
		minuteStr = "0" + minuteStr
	}
	if minute.Len > 0 {
		ret += minuteStr + minute.Suffix
	}

	secondStr := strconv.Itoa(timeNow.Second())
	if len(secondStr) == 1 {
		secondStr = "0" + secondStr
	}
	if second.Len > 0 {
		ret += secondStr + second.Suffix
	}


	return ret
}

// 传入时间并格式化成rule格式
func FTimestampsByRule(t time.Time, rule string) string {

}
