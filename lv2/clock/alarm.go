/*******
* @Author:qingmeng
* @Description:
* @File:alarm
* @Date2021/11/13
 */

package clock

import (
	"fmt"
	"time"
)

type AlarmClock struct {
	Id     int
	Key    bool //设置开关键
	Repeat bool //设置模式（是否重复）
	NextRing bool//设置下次是否响铃
}

func (a AlarmClock) Rings(value string, d time.Duration) {
	ticker:=time.NewTicker(d)
	for{
		select {
		case <-ticker.C:
			fmt.Println(value)
		}
	}
}

func (a AlarmClock) Ring(value string, d time.Duration) {
	if a.NextRing{
		d=d+24*time.Hour
	}
	timer:=time.NewTimer(d)
	for{
		<-timer.C
		fmt.Println(value)
		if a.Repeat {
			timer.Reset(24*time.Hour)
		}
	}
}

func (a AlarmClock) Timing(value string, h, m,s int) {
	//启动该闹钟的时间
	now:=time.Now()
	//设置的闹钟时间
	t:=time.Date(now.Year(),now.Month(),now.Day(),h,m,s,0,time.Local)
	if t.Before(now){
		t=t.Add(24*time.Hour)
	}
	d:=t.Sub(now)
	if a.Key{
		a.Ring(value,d)
	}
}

type Alarm interface {
	Timing(value string,h,m,s int )    //设置定时器
	Ring(value string,d time.Duration) //设置响铃
	Rings(value string,d time.Duration)//设置循环响铃
}
