/*******
* @Author:qingmeng
* @Description:
* @File:main
* @Date2021/11/13
 */

package main

import (
	"homework4/lv2/clock"
	"time"
)

func main() {
	a:=clock.AlarmClock{
		Id: -1, Key: true, Repeat: true,NextRing: true,
	}
	//基本闹钟
	go a.Timing("谁能比我卷！",2,0,0)
	go a.Timing("早八算什么，早六才是吾辈应起之时",6,0,0)
	go a.Rings("芜湖！起飞！",30*time.Second)
	clock.Start()

}
