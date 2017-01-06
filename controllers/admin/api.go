package admin

import "ecms/models"

type ApiController struct {
	baseController
}

func (this *ApiController) ShowAPIVersion() {
	this.Data["json"] = map[string]interface{}{"version":"v1"}
	this.ServeJSON()
}

func (this *ApiController) Menus() {
	var (
		menu models.Menu
		list []*models.Menu
		out *models.ApiMenus = new(models.ApiMenus)
	)

	query := menu.Query()
	count,_ := query.Count()

	if count > 0 {
		query.OrderBy("-Id").All(&list)
	}

	out.List = list
	this.Data["json"] = out
	this.ServeJSON()
}