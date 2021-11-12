/*******
* @Author:qingmeng
* @Description:
* @File:test
* @Date2021/11/12
 */

package main

import (
	"fmt"
	"strconv"
	"time"
)
func myRecover(){
	if err:=recover();err!=nil{
		fmt.Println("my recover catch:",err)
	}
}
func NewPanic() {
	defer myRecover()
	panic("test panic")
}

func main() {

	//测试错误后分协程结束
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				fmt.Println("分协程错误！")
			}
		}()
		var m map[string]string
		m["ddz"] = "yyds"
		 fmt.Println(m["ddz"])
		fmt.Println("我还想继续运行")
	}()

	time.Sleep(2*time.Second)
	fmt.Println("主协程继续运行着")

	NewPanic()
	fmt.Println("主协程继续运行着，来测试字符串转换")

	s:="123"
	s1,_:=strconv.ParseInt(s,10,0)
	fmt.Println("字符\"123\"可以计算了！", s1+s1)

	//测试时间
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format( "Mon Jan _2 15:04:05 2006"))

	//测试Parse 根据模板来解析一个格式化的字符串并返回它所代表的时间值
	fmt.Println(time.Parse("Mon Jan _2 15:04:05 2006 ","Fri Nov 12 19:43:51 2021"))

}
