package models

import "github.com/astaxie/beego/orm"

type Menu struct {
	Id       uint32
	OpenType string
	Name     string `orm:"size(100)"`
	Url      string `orm:"size(100)"`
	Icon     string `orm:"size(100)"`
	Description string `orm:"type(text);null"`
}

func (m *Menu) TableName() string {
	return "menus"
}

func (m *Menu) Query() orm.QuerySeter{
	return orm.NewOrm().QueryTable(m)
}

//添加菜单项
func (m *Menu) Insert() error{
	if _,err:= orm.NewOrm().Insert(m);err!=nil {
		return err
	}

	return nil
}

func (m *Menu) Update(fields ...string) error{
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}