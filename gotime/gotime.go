package gotime

import (
	"strconv"
	"strings"
	"time"
	"unicode"
)

type goTime struct {
	Location *time.Location `json:"location"`
}

type singleFormat struct {
	Len    int    `json:"len"`
	Suffix string `json:"suffix"`
	Index  int    `json:"index"`
}

var monthIntMap = map[string]string{
	"January":   "01",
	"February":  "02",
	"March":     "03",
	"April":     "04",
	"May":       "05",
	"June":      "06",
	"July":      "07",
	"August":    "08",
	"September": "09",
	"October":   "10",
	"November":  "11",
	"December":  "12",
}

// 时区默认是Asia/Shanghai
func NewGoTime(v interface{}) *goTime {
	var gt goTime
	if v, ok := v.(string); ok {
		if v == "" {
			v = "Asia/Shanghai"
		}
		gt.Location, _ = time.LoadLocation(v)
	} else {
		panic("please use string type for NewGoTime func")
	}
	return &gt
}

// 返回当前秒级别时间戳
func (t *goTime) Timestamps() int64 {
	timestamps := time.Now().In(t.Location)
	return timestamps.Unix()
}

// 返回当前纳秒时间
func (t *goTime) Nanosecond() int {
	timestamps := time.Now().In(t.Location)
	return timestamps.Nanosecond()
}

// 返回当前秒加纳秒的时间戳
func (t *goTime) TimestampsWithNano() string {
	timestamps := t.Timestamps()
	nanosecond := t.Nanosecond()
	timestampsWithNano := strconv.FormatInt(timestamps, 10) + "." + strconv.Itoa(nanosecond)
	return timestampsWithNano
}

// 将时间字符串转换为时间戳
func (t *goTime) RTimestamps(fTime string) int64 {
	ft, err := time.ParseInLocation("2006-01-02 15:04:05", fTime, t.Location)
	if err != nil {
		panic(err.Error())
	}

	return ft.Unix()
}

// 格式化当前时间，默认的format格式是YYYY-MM-DD HH:MM:SS
func (t *goTime) FCurrDefault() string {
	return time.Now().In(t.Location).Format("2006-01-02 15:04:05")
}

// 按照给定的格式化当前时间，例如：y-m-d h:i:s
func (t *goTime) FCorrByRule(rule string) string {
	ret := formatTime(time.Now().In(t.Location), rule)
	return ret
}

// 统一格式化方法
func formatTime(timeNow time.Time, rule string) string {
	var year singleFormat
	var month singleFormat
	var day singleFormat
	var hour singleFormat
	var minute singleFormat
	var second singleFormat
	ruleArr := strings.Split(rule, "")

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
			if key-year.Index == 1 {
				year.Suffix = ruleArr[key]
			} else if key-month.Index == 1 {
				month.Suffix = ruleArr[key]
			} else if key-day.Index == 1 {
				day.Suffix = ruleArr[key]
			} else if key-hour.Index == 1 {
				hour.Suffix = ruleArr[key]
			} else if key-minute.Index == 1 {
				minute.Suffix = ruleArr[key]
			} else if key-second.Index == 1 {
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

// 传入时间类型并格式化成rule格式
func (t *goTime) FByRule(ts time.Time, rule string) string {
	ts = ts.In(t.Location)
	ret := formatTime(ts, rule)
	return ret
}

// 传入秒级别的时间戳，转换成rule格式
func (t *goTime) FTimestampsByRule(timestamps int64, rule string) string {
	ts := time.Unix(timestamps, 0).In(t.Location)
	ret := formatTime(ts, rule)
	return ret
}

// 计算出给定时间相加时间之后的时间,并按照给定格式返回;
// 支持的时间单位: seconds,minutes,hours,days,months,years
func (t *goTime) Shift(ts time.Time, timeUnit string, rule string) string {
	fTime := ""
	ts = ts.In(t.Location)
	if strings.Contains(timeUnit, "seconds") {
		splitArr := strings.Split(timeUnit, "seconds")
		num, _ := strconv.Atoi(splitArr[0])
		fTime = formatTime(ts.Add(time.Duration(num)*time.Second), rule)
	} else if strings.Contains(timeUnit, "minutes") {
		splitArr := strings.Split(timeUnit, "minutes")
		num, _ := strconv.Atoi(splitArr[0])
		fTime = formatTime(ts.Add(time.Duration(num)*time.Minute), rule)
	} else if strings.Contains(timeUnit, "hours") {
		splitArr := strings.Split(timeUnit, "hours")
		num, _ := strconv.Atoi(splitArr[0])
		fTime = formatTime(ts.Add(time.Duration(num)*time.Hour), rule)
	} else if strings.Contains(timeUnit, "days") {
		splitArr := strings.Split(timeUnit, "days")
		num, _ := strconv.Atoi(splitArr[0])
		fTime = formatTime(ts.AddDate(0, 0, num), rule)
	} else if strings.Contains(timeUnit, "months") {
		splitArr := strings.Split(timeUnit, "months")
		num, _ := strconv.Atoi(splitArr[0])
		fTime = formatTime(ts.AddDate(0, num, 0), rule)
	} else if strings.Contains(timeUnit, "years") {
		splitArr := strings.Split(timeUnit, "years")
		num, _ := strconv.Atoi(splitArr[0])
		fTime = formatTime(ts.AddDate(num, 0, 0), rule)
	} else {
		panic("invalid time unit")
	}

	return fTime
}

// 时间sleep给定的几秒钟
func (t *goTime) Sleep(second int) {
	time.Sleep(time.Duration(second) * time.Second)
}
