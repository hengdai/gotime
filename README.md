# gotime

#### 将golang的常用时间方法进行二次封装，使用更加顺手，python版本时间模块[pytime](https://github.com/hengdai/pytime)：

示例：
```
import "gotime/gotime"

func main() {
    // 入参为时区，传入空字符串默认为 Asia/Shanghai
    gt := gotime.NewGoTime("")
    // 返回当前时间戳
    timestamps := gt.Timestamps()
    
    fmt.Println(timestamps)
}
```
输出

```
1584523430
```

所有如下方法：
```
// 返回当前秒级别时间戳
func (t *goTime) Timestamps() int64
```

```
// 返回当前纳秒时间
func (t *goTime) Nanosecond() int
```

```
// 返回当前秒加纳秒的时间戳
func (t *goTime) TimestampsWithNano() string
```

```
// 将时间字符串转换为时间戳
func (t *goTime) RTimestamps(fTime string) int64
```

```
// 格式化当前时间，默认的format格式是YYYY-MM-DD HH:MM:SS
func (t *goTime) FCurrDefault() string
```

```
// 按照给定的格式化当前时间，例如：y-m-d h:i:s
func (t *goTime) FCorrByRule(rule string) string
```

```
// 传入时间类型并格式化成rule格式
func (t *goTime) FByRule(ts time.Time, rule string) string 
```

```
// 传入秒级别的时间戳，转换成rule格式
func (t *goTime) FTimestampsByRule(timestamps int64, rule string) string
```

```
// 计算出给定时间相加时间之后的时间,并按照给定格式返回;
// 支持的时间单位: seconds,minutes,hours,days,months,years
func (t *goTime) Shift(ts time.Time, timeUnit string, rule string) string
```

```
// 时间sleep给定的几秒钟
func (t *goTime) Sleep(second int)
```
