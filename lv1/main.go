/*******
* @Author:qingmeng
* @Description:
* @File:main
* @Date2021/11/12
 */

package main

import "fmt"

func main(){
	var skillName string
	var choose int
	fmt.Println("")
	fmt.Println("选择您需要的模板：\n0.经典释放技能\n1.技能声音特效\n2.超必杀技模板")
	fmt.Scan(&choose)

	fmt.Println("释放你的技能吧:")
	fmt.Scan(&skillName)
	switch choose {
	case 0:
		ReleaseSkill(skillName, func(skillName string) {
			fmt.Println("尝尝我的厉害吧!",skillName)
		})
	case 1:
		ReleaseSkill(skillName, func(skillName string) {
			fmt.Println(skillName,"Bong~Bang")
		})
	case 2:
		ReleaseSkill(skillName, func(skillName string) {
			fmt.Println("尝尝超必杀技的威力吧！",skillName,"无人生还！")
		})

	}
}

func ReleaseSkill(skillNames string, releaseSkillFunc func(string2 string)) {
	releaseSkillFunc(skillNames)
}
