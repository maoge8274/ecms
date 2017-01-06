package main

import (
	_ "ecms/routers"
	_ "ecms/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

