package models

import (
    "fmt"
    "time"
)

/*
   代码模块化，将公共方法从main.go中抽离
*/

// UnixToTime 时间戳转换成日期字符串
func UnixToTime(timestamp int) string {
    fmt.Printf("UnixToTime-时间戳为：%v\n", timestamp)
    t := time.Unix(int64(timestamp), 0)
    return t.Format("2006-01-02 15:04:05")
}

// DateToUnix 日期转换成时间戳 2006-01-02 15:04:05
func DateToUnix(str string) int64 {
    template := "2006-01-02 15:04:05"
    t, err := time.ParseInLocation(template, str, time.Local)
    if err != nil {
        return 0
    }
    return t.Unix()
}

// GetUnix 获取当前时间戳（秒级）
func GetUnix() int64 {
    return time.Now().Unix()
}

// GetData 获取当前日期
func GetData() string {
    template := "2006-01-02 15:04:05"
    return time.Now().Format(template)
}

// GetDay 获取当前年月日
func GetDay() string {
    template := "20060102"
    return time.Now().Format(template)
}
