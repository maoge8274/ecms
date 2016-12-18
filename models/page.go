package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//页面、独立页面模型

type PageInfo struct {
	Id          int64
	Name        string `orm:"size(200)"`
	Content     string `orm:"size(10000)"`
	Title       string `orm:"size(300)"` //SEO标题
	Description string `orm:"size(500)"` //SEO页面说明
	keywords    string `orm:"size(200)"` //SEO关键字
	Status      int64                    //状态 -1 前台不显示 0 正常
	Views       int64                    //浏览量
	Addtime     time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *PageInfo) TableName() string {
	return "page_info"
}

func (m *PageInfo)Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}