package controllers

import (
	"xuhaochen/models"
)

type MenuController struct {
	BaseController
}

/**
获取左侧导航栏数据
 */
func (self *MenuController) GetMenu() {
	res, num := models.FindMenu()
	self.ajaxList("列表返回数据", 1, num, res)
}
