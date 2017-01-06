package models

import "github.com/astaxie/beego/orm"

type FaqCategory struct {
	Id uint32
}

func (m *FaqCategory) TableName() string {
	return "faq_categories"
}

func (m *FaqCategory) Query() orm.QuerySeter{
	return orm.NewOrm().QueryTable(m)
}

func (m *FaqCategory) Insert() error{
	if _,err:=orm.NewOrm().Insert(m);err!=nil{
		return err
	}
	return nil
}