package admin

import (
	"strings"
	"ecms/models"
)

//import "ecms/models"

type MenuController struct {
	baseController
}

func (this *MenuController) Get() {
	//var (
	//	menu *models.Menu
	//)
	//this.ServeJSON()
}

//添加菜单项
func (this *MenuController) Post(){
	var (
		id uint32
		name string
		url string
		icon string
		description string

		menu models.Menu
		err error
	)

	id,_ = this.GetUint32("id")

	name = strings.TrimSpace(this.GetString("name"))
	url = strings.TrimSpace(this.GetString("url"))
	icon = strings.TrimSpace(this.GetString("icon"))
	description = strings.TrimSpace(this.GetString("description"))

	menu.Icon = icon
	menu.Name = name
	menu.Url = url
	menu.Description = description

	if id>0 {
		menu.Id = id
		err = menu.Update()
	}else {
		err = menu.Insert()
	}

	if err!=nil {
		this.Data["json"] = map[string]interface{}{"code":0,"msg":"失败"}

	}else {
		this.Data["json"] = map[string]interface{}{"code":1,"msg":"成功"}
	}

	this.ServeJSON()
}