/*******
* @Author:qingmeng
* @Description:
* @File:main
* @Date2021/11/13
 */

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type User struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

type userInfo map[string]string

type Checker struct {
	userInfo userInfo
}

type Check interface {
	SignUp()//注册
	SignIn()//登陆
	Save(user User)//保存
}

func (c *Checker) SignUp() {
	fmt.Println("请输入用户名：")
	var name,psw string
	fmt.Scan(&name)
	if _,ok:=c.userInfo[name];ok{
		fmt.Println("用户已存在")
		return
	}
	fmt.Println("请输入密码:")
	for{
		fmt.Scan(&psw)
		if len(psw)>=6{
			break
		}
		fmt.Println("密码长度需大于6位，请重新输入")
	}
	for{
		fmt.Println("请确认密码")
		var again string
		fmt.Scan(&again)
		if again==psw{
			break
		}
		fmt.Println("两次密码不一致，请重新输入")
	}
	U:=User{
		UserName: name,
		PassWord: psw,
	}
	 c.Save(U)
	fmt.Println("注册成功")
}

func (c *Checker) SignIn() {
	fmt.Println("请输入用户名和密码")
	var name,psw string
	fmt.Scan(&name,psw)
	if _,ok:=c.userInfo[name];!ok{
		fmt.Println("用户名不存在")
		return
	}
	if c.userInfo[name]!=psw{
		fmt.Println("用户名错误")
		return
	}
	fmt.Println("登陆成功")
}

func (c *Checker) Save(user User) {
	f,err:=os.Open("lv3/users.data")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer f.Close()
	writer:=bufio.NewWriter(f)
	res,err:=json.Marshal(user)
	if err!=nil {
		fmt.Println(err)
	}
	n,err:=writer.WriteString(string(res))

	if err!=nil{
		fmt.Println(n,err)
	}
	writer.Flush()
	return
}

func initUsers()(userInfo,error)  {
	f,err:=os.Open("lv3/users.data")
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}
	defer f.Close()
	info:=make(userInfo)
	reader:=bufio.NewReader(f)
	for{
		buf,err:=reader.ReadBytes('\n')
		if err!=nil{
			if err==io.EOF{
				break
			}
			fmt.Println(err)
			return nil,err
		}
		var user User
		err=json.Unmarshal(buf,&user)
		if err!=nil{
			fmt.Println(err)
			continue
		}
		info[user.UserName]=user.PassWord
	}
	return info,nil
}

func main() {
	Checker:=Checker{}
	var err error
	Checker.userInfo,err=initUsers()
	if err!=nil{
		return
	}
	var op int
	for{
		fmt.Println("请输入您的操作\n1.注册\n2.登陆\n3.退出")
		fmt.Scan(&op)
		switch op {
		case 1:
			Checker.SignUp()
		case 2:
			Checker.SignIn()
		case 3:
			return
		default:
			break
		}
	}
}