package routers

import (
	"ecms/controllers"
	"ecms/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/admin", &admin.IndexController{})

	//RESTful API
	apiNs :=
		beego.NewNamespace("/api",
			beego.NSRouter("/version", &admin.IndexController{}, "get:ShowAPIVersion"),
		)

	beego.AddNamespace(apiNs)
}
