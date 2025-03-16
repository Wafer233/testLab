package time

import (
	"fmt"
	"time"
)

func Time() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	//02d输出的整数不足两位 用0补足
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func TimeStamp() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

func TimeParse() {
	t, err := time.Parse("2006-01-02 15:04:05", "2022-07-28 18:06:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", now.Format("2006/01/02 15:04:05"), loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
}

func TimeFormat() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分05秒
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

// Use method Unix to make a unix time stamp
// while use time.Unix func to trans stamp back to time
func TimeStamp2() {
	timestamp := time.Now().Unix()
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func CalTime() {
	now := time.Now()

	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println("later: ", later)

	ret := later.Sub(now)
	fmt.Println("ret: ", ret)

	res := now.Before(later)
	fmt.Println("now before later: ", res)

	res = now.After(later)
	fmt.Println("now after later: ", res)

	res = now.Equal(later)
	fmt.Println("now equal later: ", res)

}
