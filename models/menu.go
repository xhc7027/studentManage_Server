package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Menu struct {
	Id          int
	MId         int
	MName       string
	MCreateTime time.Time
	MUpdateTime time.Time
	MFid        int
	IsDelete    int
}

func (a *Menu) TableName() string {
	return TableName("x_menu")
}

func FindMenu() ([]*Menu, int64) {
	//var list []*Menu
	list := make([]*Menu,0)
	num, err := orm.NewOrm().QueryTable(TableName("x_menu")).Filter("is_delete", 0).All(&list)
	if err != nil {
		return nil, num
	}
	return list, num
}
