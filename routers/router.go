package routers

import (
	"github.com/astaxie/beego"
	"xuhaochen/controllers"
)

func init() {
	//beego.NSRouter("/*",&v1.BaseController{},"options:Options")
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.LoginController{},"*:Login")
}



