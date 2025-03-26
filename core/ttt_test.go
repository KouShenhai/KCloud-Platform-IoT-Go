package core

import (
	"fmt"
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	now := time.Now()

	// =======================格式化=======================

	// 类似输出 => 1743003091
	timestamp := now.Unix()
	fmt.Println("当前时间【秒级】：", timestamp)

	// 类似输出 => 1743003091022
	timestamp = now.UnixMilli()
	fmt.Println("当前时间【毫秒级】：", timestamp)

	// 类似输出 => 1743003091022934300
	timestamp = now.UnixNano()
	fmt.Println("当前时间【纳秒级】：", timestamp)

	// 类似输出 => 2025-03-26 23:31:31
	dateTimeStr := now.Format("2006-01-02 15:04:05")
	fmt.Println("当前时间【格式化字符串】：", dateTimeStr)

	// 类似输出 => 2025-03-26
	dateStr := now.Format("2006-01-02")
	fmt.Println("当前日期【格式化字符串】：", dateStr)

	// 类似输出 => 23:31:31
	timeStr := now.Format("15:04:05")
	fmt.Println("当前时间【格式化字符串】：", timeStr)

	// 输出 => 2025-03-26 23:29:45 +0800 CST
	timestamp = int64(1743002985)
	dateTime := time.Unix(timestamp, 0)
	fmt.Println("时间：", dateTime)

	// 类似输出 => 2025年03月26日 23:39
	dateTimeStr = now.Format("2006年01月02日 15:04")
	fmt.Println("当前时间【格式化字符串】：", dateTimeStr)

	// 类似输出 => 2025-03-26T23:40:08+08:00
	fmt.Println("时间【格式化字符串】：", now.Format(time.RFC3339))

	// 输出 => 2025-03-26 23:47:31 +0000 UTC
	dateTimeStr = "2025-03-26 23:47:31"
	dateTime, _ = time.Parse("2006-01-02 15:04:05", dateTimeStr)
	fmt.Println("时间：", dateTime)

	// =======================格式化=======================

	// =======================时区=======================

	// 输出 => Local
	fmt.Println("时区：", now.Location())

	// 输出 => 2025-03-26 23:58:31 +0800 CST
	// 设置时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	dateTime = time.Date(2025, 3, 26, 23, 58, 31, 0, loc)
	fmt.Println("时间：", dateTime)

	// 转换时区
	utc := now.UTC()
	// 类似输出 => 2025-03-26 16:01:47.8120287 +0000 UTC
	fmt.Println("UTC时间：", utc)
	dateTime = now.In(time.FixedZone("CST", 8*3600))
	// 类似输出 => 2025-03-27 00:01:47.8120287 +0800 CST
	fmt.Println("CST时间：", dateTime)

	// =======================时区=======================

	// =======================时间加减=======================

	// 类似输出 => 2025-03-28 00:03:32.952297 +0800 CST m=+86400.008917201
	add := now.Add(time.Hour * 24)
	fmt.Println("时间加1天：", add)

	// 类似输出 => 2025-03-26 00:04:09.7702908 +0800 CST m=-86399.992051199
	sub := now.Add(-time.Hour * 24)
	fmt.Println("时间减1天：", sub)

	// 类似输出 => 24h0m0s
	fmt.Println("时间差：", now.Sub(sub))

	// 类似输出 => 24
	fmt.Println("时间差：", now.Sub(sub).Hours())

	// 类似输出 => 2025-04-28 00:06:14.4762206 +0800 CST
	add = now.AddDate(0, 1, 1)
	fmt.Println("时间加1月1天：", add)

	t1 := now
	t2 := now.Add(time.Hour * 1)
	// 输出 => true
	fmt.Println("比较：", t2.After(t1))
	// 输出 => true
	fmt.Println("比较：", t1.Before(t2))

	// =======================时间加减=======================

	// =======================其他=======================

	// 类似输出 => 2025
	year := now.Year()
	fmt.Println("年：", year)

	// 类似输出 => 3
	month := now.Month()
	fmt.Println("月：", month)

	// 类似输出 => 27
	day := now.Day()
	fmt.Println("日：", day)

	// 类似输出 => 0
	hour := now.Hour()
	fmt.Println("时：", hour)

	// 类似输出 => 15
	minute := now.Minute()
	fmt.Println("分：", minute)

	// 类似输出 => 54
	second := now.Second()
	fmt.Println("秒：", second)

	// 类似输出 => Thursday
	weekday := now.Weekday()
	fmt.Println("星期：", weekday)

	// 类似输出 => 2025-03-26 08:00:00 +0800 CST
	truncate := now.Truncate(24 * time.Hour)
	fmt.Println("时间截断：", truncate)

	// 类似输出 => 2025-03-27 00:00:00 +0800 CST
	round := now.Round(time.Hour)
	fmt.Println("时间四舍五入：", round)

	// =======================其他=======================

	// =======================周期任务=======================

	// 等待2秒【定时器】
	timer := time.NewTimer(2 * time.Second)
	<-timer.C

	// 每秒执行一次【周期任务】
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		fmt.Println("每秒执行一次")
	}

	// =======================周期任务=======================

}
