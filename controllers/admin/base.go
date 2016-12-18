package admin

import "github.com/astaxie/beego"

type baseController struct {
	beego.Controller
	userid         int64
	username       string
	nickname       string
	controllerName string
	actionName     string
	admindir       string
}