package models

import "github.com/astaxie/beego/orm"

type Admin struct {
	Id        int
	LoginName string
	PassWord  string
	Status    int
}

func (a *Admin) TableName() string {
	return TableName("x_admin")
}

func AdminGetByName(loginName string) (*Admin, error) {
	a := new(Admin)
	err := orm.NewOrm().QueryTable(TableName("x_admin")).Filter("login_name", loginName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
