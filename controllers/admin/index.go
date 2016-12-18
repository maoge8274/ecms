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

