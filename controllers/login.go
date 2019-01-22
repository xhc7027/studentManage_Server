package controllers

import (
	"fmt"
	"xuhaochen/models"
)

type LoginController struct {
	BaseController
}
type User struct {
	UserName string
	PassWord  string
}

func (self *LoginController) Login() {
	var user User
	user = self.AddUser()
	fmt.Println(user.UserName)
	if self.isPost() {
		fmt.Println(self.Input())
		userName := user.UserName
		passWord := user.PassWord

		if userName != "" && passWord != "" {
			user, err := models.AdminGetByName(userName)
			fmt.Println(user)
			if err != nil || user.PassWord != passWord {
				self.ajaxMsg("账号或密码错误", MSG_ERR)
			} else if user.Status == -1 {
				self.ajaxMsg("该账号已禁用", MSG_ERR)
			} else {
				self.ajaxMsg("登陆成功", MSG_OK)
			}
		}
	}
	self.ajaxMsg("请求方式错误", MSG_ERR)
}
