/*******
* @Author:qingmeng
* @Description:
* @File:start
* @Date2021/11/13
 */

package clock

import "fmt"

func Start() {
	var alarms []AlarmClock
	for  {
		fmt.Println("哎，start里面的功能不用试了，基本lv2的还在，想做做ex，感觉思维不对，想简单了，就做了一半没做了，不该用切片的方法，肯定不是想考察的点")
		fmt.Println("1.增加闹钟\n2.删除闹钟\n3.取消某闹钟的下次提醒\n0.退出程序")
		fmt.Println("请输入您需要的功能")
		op:=0
		fmt.Scan(&op)
		switch op {
		case 0:
			fmt.Println("程序已安全退出")
			return
		case 1:
			fmt.Println("请设置您的闹钟（编号，是否开启，是否重复，下次是否提醒）")
			var(
				id int
				key bool
				repeat bool
				next bool
			)
			fmt.Scan(&id,&key,&repeat,&next)
			alarm:=AlarmClock{
			id,key,repeat,next,
			}
			fmt.Println("请输入提醒内容和提醒的时间(h,m,s)")
			var(
				value string
				h,m,s int
			)
			fmt.Scanln(&value)
			fmt.Scan(&h,&m,&s)
			fmt.Println("闹钟设置成功")
			go alarm.Timing(value,h,m,s)
			alarms=append(alarms, alarm)
		case 2:
			fmt.Println("您已经设的闹钟有:")
				fmt.Println(alarms)
		case 3:
			break
		default:
			break
		}
	}
}
