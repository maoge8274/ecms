package routers

import (
	"ecms/controllers"
	"github.com/astaxie/beego"
	"ecms/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/admin", &admin.IndexController{})
}
