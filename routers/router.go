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
			beego.NSRouter("/version", &admin.ApiController{}, "get:ShowAPIVersion"),
			beego.NSRouter("/menus",&admin.ApiController{},"get:Menus"),
			beego.NSRouter("/menu",&admin.MenuController{}),
			//beego.NSRouter("/menu/:id",&admin.MenuController{}),
		)

	beego.AddNamespace(apiNs)
}
