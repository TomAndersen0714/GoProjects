package time

import (
	"fmt"
	"time"
)

const (
	TimerFormat = "2006-01-02 15:04:05.000"
)

func Demo() {
	dialogTime := "2022-03-25 23:59:56.931"
	msgTime := "2022-03-25 15:59:57.341"
	localTime, _ := time.ParseInLocation("2006-01-02 15:04:05", dialogTime, time.Local)

	fmt.Println(localTime)
	fmt.Println(localTime.Unix())

	parseMsgTime, _ := time.Parse(TimerFormat, msgTime)
	fmt.Println(parseMsgTime)
	fmt.Println(parseMsgTime.UnixMicro())
}
