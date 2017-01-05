package admin

type IndexController struct {
	baseController
}

type LoginController struct {
	baseController
}


func (this *IndexController) Get(){
	this.TplName = "admin/index.html"
}

func (this *IndexController) ShowAPIVersion() {
	this.Data["json"] = map[string]interface{}{"version":"v1"}
	this.ServeJSON()
}