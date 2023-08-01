package router

import (
	"github.com/astaxie/beego"
	"github.com/wangpan-hqu/beego_use/controller"
)

func init() {
	beego.Router("/api/signin", &controller.ApiController{}, "POST:Signin")
}
